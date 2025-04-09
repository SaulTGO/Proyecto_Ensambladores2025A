package main

import (
	"bufio"
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
	tmpl.ExecuteTemplate(w, "fase2.html", nil)
}

// ServeFase3 Ejecutar fase3.html al acceder a /Fase3
func ServeFase3(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "fase3.html", nil)
}

// Ejecutar al subir un archivo
func handleUploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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
		defer file.Close()

		// Crear un scanner para leer línea por línea
		scanner := bufio.NewScanner(file)

		//Guardar cada linea del archivo
		for i := 0; scanner.Scan(); i++ {
			//Llenar la variable con las lineas de codigo

			SourceCode = append(SourceCode, scanner.Text())
		}

		//Verificar que el escaner haga su chamba
		if err := scanner.Err(); err != nil {
			http.Error(w, "Error al leer el archivo: "+err.Error(), http.StatusInternalServerError)
			return
		}

		//Una vez hechas las validaciones, despleagar el contenido
		//tmpl.ExecuteTemplate(w, "display_source_code.html", SourceCode)

		AnalizeSourceCode(SourceCode)

		//tmpl.ExecuteTemplate(w, "display_parsed_code.html", InstructionsSlice)

		http.Redirect(w, r, "/Fase1", http.StatusSeeOther)
		return
	}
}
