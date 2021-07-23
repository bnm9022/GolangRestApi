
## Testing
Use go run main.go to run the server

curl -X GET http://localhost:8000/applicationMDs
curl -X GET http://localhost:8000/applicationMD/0
curl -X POST -d "C:\Users\kaiyli\Desktop\temp\testdata\valid1.yaml" http://localhost:8000/applicationMD
curl -X POST -d "C:\Users\kaiyli\Desktop\temp\testdata\valid2.yaml" http://localhost:8000/applicationMD
curl -X POST -d "C:\Users\kaiyli\Desktop\temp\testdata\invalid1.yaml" http://localhost:8000/applicationMD
curl -X POST -d "C:\Users\kaiyli\Desktop\temp\testdata\invalid2.yaml" http://localhost:8000/applicationMD
curl -X GET http://localhost:8000/applicationMDs
curl -X GET "http://localhost:8000/applicationMD?title=Valid%20App%201"
curl -X GET "http://localhost:8000/applicationMD?title=Valid%20App%201&&version=1.0.1"
curl -X PUT -d "C:\Users\kaiyli\Desktop\temp\testdata\valid3.yaml" "http://localhost:8000/applicationMD?title=Valid%20App%202"
curl -X GET http://localhost:8000/applicationMDs
curl -X DELETE "http://localhost:8000/applicationMD?title=Valid%20App%202"

Run tests:

C:\Users\kaiyli\Desktop\temp>go test -v
=== RUN   TestGetApplicationMD
--- PASS: TestGetApplicationMD (0.00s)
=== RUN   TestPostAndDeleteApplicationMD
--- PASS: TestPostAndDeleteApplicationMD (0.00s)
=== RUN   TestPostAndGetApplicationMD
--- PASS: TestPostAndGetApplicationMD (0.00s)
=== RUN   TestPostAndGetByParametersApplicationMD
--- PASS: TestPostAndGetByParametersApplicationMD (0.00s)
=== RUN   TestPutAndGetApplicationMD
--- PASS: TestPutAndGetApplicationMD (0.00s)
PASS
ok      temp    0.163s

## Curl Output
C:\Users\kaiyli\Desktop\curl>curl -X GET http://localhost:8000/applicationMDs
[]

C:\Users\kaiyli\Desktop\curl>curl -X GET http://localhost:8000/applicationMD/0
Record not find
C:\Users\kaiyli\Desktop\curl>curl -X POST -d "C:\Users\kaiyli\Desktop\temp\testdata\valid1.yaml" http://localhost:8000/applicationMD
title: Valid App 1
version: 0.0.1
maintainers:
- name: firstmaintainer app1
  email: firstmaintainer@hotmail.com
- name: secondmaintainer app1
  email: secondmaintainer@gmail.com
company: Random Inc.
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: |-
  ### Interesting Title
  Some application content, and description

C:\Users\kaiyli\Desktop\curl>curl -X POST -d "C:\Users\kaiyli\Desktop\temp\testdata\valid2.yaml" http://localhost:8000/applicationMD
title: Valid App 2
version: 1.0.1
maintainers:
- name: AppTwo Maintainer
  email: apptwo@hotmail.com
company: Upbound Inc.
website: https://upbound.io
source: https://github.com/upbound/repo
license: Apache-2.0
description: |-
  ### Why app 2 is the best
  Because it simply is...

C:\Users\kaiyli\Desktop\curl>curl -X POST -d "C:\Users\kaiyli\Desktop\temp\testdata\invalid1.yaml" http://localhost:8000/applicationMD
Maintainer email is not in correct format
C:\Users\kaiyli\Desktop\curl>curl -X POST -d "C:\Users\kaiyli\Desktop\temp\testdata\invalid2.yaml" http://localhost:8000/applicationMD
Version could not be empty
C:\Users\kaiyli\Desktop\curl>curl -X GET http://localhost:8000/applicationMDs
- title: Valid App 1
  version: 0.0.1
  maintainers:
  - name: firstmaintainer app1
    email: firstmaintainer@hotmail.com
  - name: secondmaintainer app1
    email: secondmaintainer@gmail.com
  company: Random Inc.
  website: https://website.com
  source: https://github.com/random/repo
  license: Apache-2.0
  description: |-
    ### Interesting Title
    Some application content, and description
- title: Valid App 2
  version: 1.0.1
  maintainers:
  - name: AppTwo Maintainer
    email: apptwo@hotmail.com
  company: Upbound Inc.
  website: https://upbound.io
  source: https://github.com/upbound/repo
  license: Apache-2.0
  description: |-
    ### Why app 2 is the best
    Because it simply is...

C:\Users\kaiyli\Desktop\curl>curl -X GET "http://localhost:8000/applicationMD?title=Valid%20App%201"
- title: Valid App 1
  version: 0.0.1
  maintainers:
  - name: firstmaintainer app1
    email: firstmaintainer@hotmail.com
  - name: secondmaintainer app1
    email: secondmaintainer@gmail.com
  company: Random Inc.
  website: https://website.com
  source: https://github.com/random/repo
  license: Apache-2.0
  description: |-
    ### Interesting Title
    Some application content, and description

C:\Users\kaiyli\Desktop\curl>curl -X GET "http://localhost:8000/applicationMD?title=Valid%20App%201&&version=1.0.1"
Could not find based on search parameters
C:\Users\kaiyli\Desktop\curl>curl -X PUT -d "C:\Users\kaiyli\Desktop\temp\testdata\valid3.yaml" "http://localhost:8000/applicationMD?title=Valid%20App%202"
Updated the following records:title: Valid App 2
version: 1.0.1
maintainers:
- name: AppTwo Maintainer
  email: apptwo@hotmail.com
company: Upbound Inc.
website: https://upbound.io
source: https://github.com/upbound/repo
license: Apache-2.0
description: |-
  ### Why app 3 is the best
  Because it simply is...

C:\Users\kaiyli\Desktop\curl>curl -X GET http://localhost:8000/applicationMDs
- title: Valid App 1
  version: 0.0.1
  maintainers:
  - name: firstmaintainer app1
    email: firstmaintainer@hotmail.com
  - name: secondmaintainer app1
    email: secondmaintainer@gmail.com
  company: Random Inc.
  website: https://website.com
  source: https://github.com/random/repo
  license: Apache-2.0
  description: |-
    ### Interesting Title
    Some application content, and description
- title: Valid App 2
  version: 1.0.1
  maintainers:
  - name: AppTwo Maintainer
    email: apptwo@hotmail.com
  company: Upbound Inc.
  website: https://upbound.io
  source: https://github.com/upbound/repo
  license: Apache-2.0
  description: |-
    ### Why app 3 is the best
    Because it simply is...

C:\Users\kaiyli\Desktop\curl>curl -X DELETE "http://localhost:8000/applicationMD?title=Valid%20App%202"
Removed the following records:- title: Valid App 2
  version: 1.0.1
  maintainers:
  - name: AppTwo Maintainer
    email: apptwo@hotmail.com
  company: Upbound Inc.
  website: https://upbound.io
  source: https://github.com/upbound/repo
  license: Apache-2.0
  description: |-
    ### Why app 3 is the best
    Because it simply is...

C:\Users\kaiyli\Desktop\curl>curl -X GET http://localhost:8000/applicationMDs
- title: Valid App 1
  version: 0.0.1
  maintainers:
  - name: firstmaintainer app1
    email: firstmaintainer@hotmail.com
  - name: secondmaintainer app1
    email: secondmaintainer@gmail.com
  company: Random Inc.
  website: https://website.com
  source: https://github.com/random/repo
  license: Apache-2.0
  description: |-
    ### Interesting Title
    Some application content, and description

C:\Users\kaiyli\Desktop\curl>
