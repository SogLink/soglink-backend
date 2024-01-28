package models

type GetDoctorListResponse struct {
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	SpecialtyName string `json:"specialty_name"`
	Education     string `json:"education"`
	Certificates  string `json:"certificate:"`
	ClinicName    string `json:"clinic_name"`
}
