define({ "api": [
  {
    "type": "post",
    "url": "/permissionLetter",
    "title": "generate permission letter",
    "name": "permissionLetter",
    "group": "all",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "clubName",
            "description": "<p>Name of your club</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Name of your event</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "toDate",
            "description": "<p>ending date</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "fromDate",
            "description": "<p>start date</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "toTime",
            "description": "<p>start time</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "fromTime",
            "description": "<p>ending time</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "budget",
            "description": "<p>budget</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "description",
            "description": "<p>event description</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "category",
            "description": "<p>category of the event</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "venue",
            "description": "<p>event venue</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "attendance",
            "description": "<p>Name of your club</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "expectedParticipants",
            "description": "<p>Expected Participants in the event</p>"
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "facultyCoordinator",
            "description": "<p>faculty coordinator details (Participant Object)</p>"
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "studentCoordinator",
            "description": "<p>student coordinator details (Participant Object)</p>"
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "guest",
            "description": "<p>guest details (Guest object)</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "PROrequest",
            "description": "<p>PRO department requests</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "campusEngineerRequest",
            "description": "<p>Campus engineer requests</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "duration",
            "description": "<p>duration of event</p>"
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "mainSponsor",
            "description": "<p>sponsor details (Participant Object)</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n   \"clubName\":\"GDG\",\n   \"name\":\"DEVFEST\",\n   \"toDate\":\"10TH OCTOBER\",\n   \"fromDate\":\"8TH OCTOBER\",\n   \"toTime\":\"10 PM\",\n   \"fromTime\":\"11 AM\",\n   \"budget\":\"200000\",\n   \"description\":\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\",\n   \"category\":\"TECHNICAL\",\n   \"venue\":\"ANNA AUDI\",\n   \"attendance\":\"4000\",\n   \"expectedParticipants\":\"4000\",\n   \"facultyCoordinator\":{\n      \"name\":\"NOORU MAA\",\n      \"registrationNumber\":\"17BBE1010\",\n      \"email\":\"SDADAS@A.COM\",\n      \"phoneNumber\":\"919191991911\",\n      \"gender\":\"M\",\n      \"eventsAttended\":\"ALL\"\n   },\n   \"studentCoordinator\":{\n      \"name\":\"NOORU BAAP\",\n      \"registrationNumber\":\"17BBE1010\",\n      \"email\":\"SDADAS@A.COM\",\n      \"phoneNumber\":\"919191991911\",\n      \"gender\":\"M\",\n      \"eventsAttended\":\"ALL\"\n   },\n   \"guest\":{\n      \"name\":\"ALLAHH DAAS\",\n      \"email\":\"ASDSAD#ASD.COM\",\n      \"phoneNumber\":\"11111111111\",\n      \"gender\":\"F\",\n      \"stake\":\"SOME MONAYYYY\",\n      \"locationOfStay\":\"TERA GHAR\"\n   },\n   \"PROrequest\":\"SAJDOOSIJANDFSAKFDSAFD\",\n   \"campusEngineerRequest\":\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\n   \"duration\":\"16 hours\",\n   \"mainSponsor\":{\n      \"name\":\"ALLAHH DAAS\",\n      \"email\":\"ASDSAD#ASD.COM\",\n      \"phoneNumber\":\"11111111111\",\n      \"gender\":\"F\",\n      \"stake\":\"SOME MONAYYYY\",\n      \"locationOfStay\":\"TERA GHAR\"\n   }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "html",
            "optional": false,
            "field": "html",
            "description": "<p>Get html.</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./main/controller/generatePermissionLetter.go",
    "groupTitle": "all",
    "sampleRequest": [
      {
        "url": "http://localhost/api/v1/permissionLetter"
      }
    ]
  }
] });
