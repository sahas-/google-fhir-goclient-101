package main

var resourceOrder = []string{
	"Organization",
	"Location",
	"Practitioner",
	"PractitionerRole",
	"Patient",
	"Encounter",
	"Medication",
	"MedicationRequest",
	"MedicationStatement",
}

const (
	FHIR_SERVER                          = "http://localhost:8080/fhir/"
	INPUT_FOLDER                         = "/Users/sahaswaranamam/Documents/Explore/synthea/output/fhir"
	API_KEY                              = "B29CD476-119B-4928-BB5A-C6191A7B435E"
	ORG_RESOURCE                         = "Organization"
	LOCATION_RESOURCE                    = "Location"
	PRACTITIONER_RESOURCE                = "Practitioner"
	PRACTITIONER_ROLE_RESOURCE           = "PractitionerRole"
	PATIENT_RESOURCE                     = "Patient"
	CONDITION_RESOURCE                   = "Condition"
	ALLERGY_RESOURCE                     = "AllergyIntolerance"
	MEDICATION_RESOURCE                  = "Medication"
	DEVICE_RESOURCE                      = "Device"
	CAREPLAN_RESOURCE                    = "CarePlan"
	CARETEAM_RESOURCE                    = "CareTeam"
	MEDICATION_ADMINISTIONATION_RESOURCE = "MedicationAdministration"
	MEDICATION_REQ_RESOURCE              = "MedicationRequest"
	MEDICATION_STATEMENT_RESOURCE        = "MedicationStatement"
	OBSERVATION_RESOURCE                 = "Observation"
	DIAGNOSTIC_REPORT_RESOURCE           = "DiagnosticReport"
	IMMUNIZATION_RESOURCE                = "Immunization"
	ENCOUNTER_RESOURCE                   = "Encounter"
	CLAIM_RESOURCE                       = "Claim"
	EXPLANATION_OF_BENEFIT_RESOURCE      = "ExplanationOfBenefit"
	DOCUMENT_REFERENCE_RESOURCE          = "DocumentReference"
	PROVENANCE_RESOURCE                  = "Provenance"
	SUPPLYDELIVERY_RESOURCE              = "SupplyDelivery"
	IMAGINGSTUDY_RESOURCE                = "ImagingStudy"
	PROCEDURE_RESOURCE                   = "Procedure"
	timeZone                             = "America/Los_Angeles"
)
