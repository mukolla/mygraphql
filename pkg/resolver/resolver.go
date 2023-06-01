package resolver

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/graphql-go/graphql"
)

type Developer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Work struct {
	ID        string `json:"id"`
	Position  string `json:"position"`
	Title     string `json:"title"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type WorkDeveloper struct {
	DeveloperID string     `json:"developerId"`
	WorkID      string     `json:"workId"`
	Developer   *Developer `json:"developer"`
}

var Developers = []Developer{
	Developer{ID: "1", Name: "John Doe", Email: "john@example.com"},
	Developer{ID: "2", Name: "Jane Smith", Email: "jane@example.com"},
}

var works = []Work{
	Work{ID: "1", Position: "Software Engineer", Title: "Work 1", StartDate: "2022-01-01", EndDate: "2022-12-31"},
	Work{ID: "2", Position: "Frontend Developer", Title: "Work 2", StartDate: "2023-01-01", EndDate: ""},
}

var workDevelopers = []WorkDeveloper{
	WorkDeveloper{DeveloperID: "1", WorkID: "1"},
	WorkDeveloper{DeveloperID: "2", WorkID: "2"},
}

func ResolveDeveloper(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	for _, developer := range Developers {
		if developer.ID == id {
			return developer, nil
		}
	}
	return nil, nil
}

func GetDeveloperByID(id string) *Developer {
	for _, developer := range Developers {
		if developer.ID == id {
			return &developer
		}
	}
	return nil
}

func ResolveDevelopers(p graphql.ResolveParams) (interface{}, error) {
	return Developers, nil
}

func Works(p graphql.ResolveParams) (interface{}, error) {
	return works, nil
}

func ResolveWork(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	for _, work := range works {
		if work.ID == id {
			return work, nil
		}
	}
	return nil, nil
}

func ResolveWorkDeveloper(p graphql.ResolveParams) (interface{}, error) {
	developerID, _ := p.Args["developerId"].(string)
	workID, _ := p.Args["workId"].(string)
	for _, wd := range workDevelopers {
		if wd.DeveloperID == developerID && wd.WorkID == workID {
			return wd, nil
		}
	}
	return nil, nil
}

func GetWorkDeveloper(p graphql.ResolveParams) (interface{}, error) {
	workDevs := workDevelopers

	for i := range workDevs {
		developerID := workDevs[i].DeveloperID
		developer := GetDeveloperByID(developerID)
		workDevs[i].Developer = developer
	}

	return workDevs, nil
}

func getDevelopers() []Developer {
	return Developers
}

func AddDeveloper(name string, email string) {
	id := generateUniqueID() // Функція для генерації унікального ID
	developer := Developer{ID: id, Name: name, Email: email}
	Developers = append(Developers, developer)
}

func generateUniqueID() string {
	// Генеруємо випадковий байтовий масив
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// Обробка помилки при генерації випадкових байтів
		// Тут можна виконати відповідну обробку помилки, наприклад, викинути виключення
		panic(err)
	}

	// Конвертуємо байтовий масив в шістнадцятковий рядок
	uniqueID := hex.EncodeToString(randomBytes)

	return uniqueID
}

//developerID, ok := p.Info.RootValue.(map[string]interface{})["developerId"].(string)
//if !ok {
//	return nil, nil
//}
//
//fmt.Println(developerID)

//var workDevelopersWithDeveloperInfo []map[string]interface{}
//
////developerID, _ := p.Args["developerId"].(string)
//
//fmt.Println(p.Args)
//
//developer := GetDeveloperByID("1")
//fmt.Println(developer)
//
//developerInfo := map[string]interface{}{
//	"developerId": 7,
//	"workId":      6,
//	"developer":   "d",
//}
//workDevelopersWithDeveloperInfo = append(workDevelopersWithDeveloperInfo, developerInfo)
//return workDevelopersWithDeveloperInfo, nil

//for _, wd := range workDevelopers {
//
//	fmt.Println("DeveloperID: ", wd.DeveloperID)
//	fmt.Println(GetDeveloperByID(wd.DeveloperID))
//
//	dev := GetDeveloperByID(wd.DeveloperID)
//	if dev == nil {
//		continue
//	}
//	developerInfo := map[string]interface{}{
//		"developerId": wd.DeveloperID,
//		"workId":      wd.WorkID,
//		"developer":   dev,
//	}
//	workDevelopersWithDeveloperInfo = append(workDevelopersWithDeveloperInfo, developerInfo)
//}
//
//return workDevelopersWithDeveloperInfo, nil

/*Resolve: func(p graphql.ResolveParams) (interface{}, error) {

	fmt.Println(p.Info.)

	//project, ok := p.Info.ParentType.(*graphql.Object)
	//if !ok {
	//	return nil, nil
	//}

	//
	//fmt.Println(project.Fields())
	//developerID := p.Source.(map[string]interface{})["developerId"].(string)

	//project, ok := p.Info.ParentType.(*graphql.Object)
	//if !ok {
	//	return nil, nil
	//}
	//
	//fields := project.Fields()
	//developerIDField, ok := fields["developerId"]
	//if !ok {
	//	return nil, nil
	//}
	//
	//developerIDType, ok := developerIDField.Type.(*graphql.NonNull).OfType.(*graphql.Scalar)
	//if !ok {
	//	return nil, nil
	//}
	//
	//developerID := developerIDType.Name()

	// Використовуйте значення developerID за необхідність
	//fmt.Println("developerId:", developerID)

	//developerIDField := context.Background()

	//fmt.Println(developerIDField)

	developer := resolver.GetDeveloperByID("2")
	return developer, nil
},*/
