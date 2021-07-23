package main
import (
    "io/ioutil"
    "net/http"
    "gopkg.in/yaml.v2"
    "net/mail"
)


func checkSearchParameter (responseWriter http.ResponseWriter, r *http.Request) (bool, []ApplicationMD, string, string) {
    var searchResults []ApplicationMD
    title := r.URL.Query().Get("title")
    version := r.URL.Query().Get("version")
    if len(title) == 0 && len(version) == 0 {
        responseWriter.WriteHeader(404)
        responseWriter.Write([]byte(`Please add more parameters to search`))
        return false, searchResults, title, version
    }
   
    for _, applicationMD := range ApplicationMDs {
        if  (len(title) == 0 || applicationMD.Title == title) && (len(version) == 0 || applicationMD.Version == version) {
            searchResults = append(searchResults, applicationMD)
        }
    }

    if (len(searchResults) == 0) {
        responseWriter.WriteHeader(404)
        responseWriter.Write([]byte(`Could not find based on search parameters`))
        return false, searchResults, title, version
    }
    return true, searchResults, title, version
}

func createNewApplicationMDCheck (responseWriter http.ResponseWriter, r *http.Request) (bool, ApplicationMD){
    var applicationMD ApplicationMD 
    reqBody, _ := ioutil.ReadAll(r.Body)
    yamlFile, err := ioutil.ReadFile(string(reqBody))
    if err != nil {
        responseWriter.WriteHeader(400)
        responseWriter.Write([]byte(`Could not read provided yaml file path`))
        return false, applicationMD
    }
    
    err = yaml.Unmarshal(yamlFile, &applicationMD)
    if err != nil {
        responseWriter.WriteHeader(400)
        responseWriter.Write([]byte(`Unable to unmarshal given yaml data`))
        return false, applicationMD
    }

    if len(applicationMD.Title) == 0 {
        responseWriter.WriteHeader(400)
        responseWriter.Write([]byte(`Title could not be empty`))
        return false, applicationMD
    }

    if len(applicationMD.Version) == 0 {
        responseWriter.WriteHeader(400)
        responseWriter.Write([]byte(`Version could not be empty`))
        return false, applicationMD
    }

    for _, maintainer := range applicationMD.Maintainers {
        if len(maintainer.Email) == 0 {
            responseWriter.WriteHeader(400)
            responseWriter.Write([]byte(`Maintainer name could not be empty`))
            return false, applicationMD
        }
        _, err := mail.ParseAddress(maintainer.Email)
        if err != nil {
            responseWriter.WriteHeader(400)
            responseWriter.Write([]byte(`Maintainer email is not in correct format`))
            return false, applicationMD
        }
    }
    return true, applicationMD
}
