package main

import (
	"bufio"
	"mime/multipart"
	"net/http"
)

// ServeIndex /* Estos metodos ejecutan el archivo correspondiente cuando la ruta es llamada
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

// ServeFase1 Ejecutar fase1.html al acceder a /Fase1
func ServeFase1(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handleUploadFile(w, r)
		return
	}

	codeData := Data{
		SourceCode:        SourceCode,
		InstructionsSlice: InstructionsSlice,
		FileProcessed:     len(SourceCode) > 0,
	}
	err := tmpl.ExecuteTemplate(w, "fase1.html", codeData)
	if err != nil {
		return
	}

}

// ServeFase2 Ejecutar fase2.html al acceder a /Fase2
func ServeFase2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handleFileFase2(w, r)
		return
	}

	fase2Data := Fase2Data{
		SourceCode:    SourceCode,
		Validations:   LinesValidation,
		SymbolTable:   SymbolTable,
		FileProcessed: len(SourceCode) > 0,
	}

	err := tmpl.ExecuteTemplate(w, "fase2.html", fase2Data)
	if err != nil {
		return
	}
}

// ServeFase3 Ejecutar fase3.html al acceder a /Fase3
func ServeFase3(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handleFileFase3(w, r)
		return
	}

	err := tmpl.ExecuteTemplate(w, "fase3.html", nil)
	if err != nil {
		return
	}
}

// Ejecutar al subir un archivo
func handleUploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		processFile(w, r)

		//Una vez hechas las validaciones, despleagar el contenido

		InstructionsSlice = []AsmInstruction{}
		AnalizeSourceCode(SourceCode)

		http.Redirect(w, r, "/Fase1", http.StatusSeeOther)
		return
	}
}

func handleFileFase2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		processFile(w, r)

		InstructionsSlice = []AsmInstruction{}
		AnalizeSourceCode(SourceCode)

		StackLines = make(map[int]string)
		CodeLines = make(map[int]string)
		DataLines = make(map[int]string)

		SymbolTable = []Symbol{}
		LinesValidation = []LineStatus{}
		SemanticAnalysis()

		http.Redirect(w, r, "/Fase2", http.StatusSeeOther)
		return
	}
}

func handleFileFase3(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		processFile(w, r)

		http.Redirect(w, r, "/Fase3", http.StatusSeeOther)
		return
	}
}

func processFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) //Tamaño maximo
	if err != nil {
		http.Error(w, "Error al parsear el formulario: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar si existe el archivo
	files := r.MultipartForm.File["uploaded"]
	if len(files) == 0 {
		http.Error(w, "No se encontró el archivo", http.StatusBadRequest)
		return
	}

	//Abrir el archivo
	fileheader := files[0]

	//Abrir el archivo
	file, err := fileheader.Open()
	if err != nil {
		http.Error(w, "Error al abrir el archivo: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Crear un scanner para leer línea por línea
	scanner := bufio.NewScanner(file)
	SourceCode = make(map[int]string)

	//Guardar cada linea del archivo
	for i := 0; scanner.Scan(); i++ {
		//Llenar la variable con las lineas de codigo
		SourceCode[i] = scanner.Text()
	}

	//Verificar que el escaner haga su chamba
	if err := scanner.Err(); err != nil {
		http.Error(w, "Error al leer el archivo: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
