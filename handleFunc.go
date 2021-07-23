package main

import (
    "strconv"
    "net/http"
    "gopkg.in/yaml.v2"
    "github.com/gorilla/mux"
)


func returnAllApplicationMDs(responseWriter http.ResponseWriter, r *http.Request) {
    yaml.NewEncoder(responseWriter).Encode(ApplicationMDs)
}

func returnSingleApplicationMD(responseWriter http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    found := false

    for _, applicationMD := range ApplicationMDs {
        if applicationMD.Id == key {
            found = true
            yaml.NewEncoder(responseWriter).Encode(applicationMD)
            return
        }
    }
    if found == false {
        responseWriter.WriteHeader(404)
        responseWriter.Write([]byte(`Record not find`))
    }
}

func returnApplicationMDLists(responseWriter http.ResponseWriter, r *http.Request) {
    check, searchResults, _, _ := checkSearchParameter(responseWriter, r)
    if (check == false) {
        return
    }
    yaml.NewEncoder(responseWriter).Encode(searchResults)
}

func createNewApplicationMD(responseWriter http.ResponseWriter, r *http.Request) {
    newApplicationMDCheck, applicationMD := createNewApplicationMDCheck(responseWriter, r)
    if (newApplicationMDCheck == false) {
        return
    }

    yaml.NewEncoder(responseWriter).Encode(applicationMD)
    applicationMD.Id = strconv.Itoa(currentId)
    currentId = currentId + 1
    ApplicationMDs = append(ApplicationMDs, applicationMD)
}

func deleteApplicationMD(responseWriter http.ResponseWriter, r *http.Request) {
    check, searchResults, title, version := checkSearchParameter(responseWriter, r)
    if (check == false) {
        return
    }

    for index, applicationMD := range ApplicationMDs {
        if (len(title) == 0 || applicationMD.Title == title) && (len(version) == 0 || applicationMD.Version == version) {
            ApplicationMDs = append(ApplicationMDs[:index], ApplicationMDs[index+1:]...)
        }
    }
    
    responseWriter.Write([]byte(`Removed the following records:`))
    yaml.NewEncoder(responseWriter).Encode(searchResults)
}

func putApplicationMD(responseWriter http.ResponseWriter, r *http.Request) {
    check, searchResults, title, version := checkSearchParameter(responseWriter, r)
    if (check == false) {
        return
    }

    if (len(searchResults) > 1) {
        responseWriter.WriteHeader(404)
        responseWriter.Write([]byte(`There are more than one records, please change search parameters`))
        return
    }

    newApplicationMDCheck, applicationMD := createNewApplicationMDCheck(responseWriter, r)
    if (newApplicationMDCheck == false) {
        return
    }

    for index, oldApplicationMD := range ApplicationMDs {
        if (len(title) == 0 || oldApplicationMD.Title == title) && (len(version) == 0 || oldApplicationMD.Version == version) {
            ApplicationMDs[index] = applicationMD
        }
    }
    
    responseWriter.Write([]byte(`Updated the following records:`))
    yaml.NewEncoder(responseWriter).Encode(applicationMD)
}