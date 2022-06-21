package ts

var _tpl = `import axios from "axios"

{{ range .Models }}
interface {{.Name}}{
	{{range .Fields}}
		{{.Name}}:{{.Type}},
	{{end}}
}
{{ end}}

export default {
    api: {
        {{login:"",}}
    },
    async post(api:string,body?:any, header?:any) {
        return axios.post(api,body, {headers:header})
    },
    async get(api:string,body?:any, header?:any)  {
        return axios.get(api,body,{headers:header})
    },
    async login(body:LoginRequest,header?:any) {
        return this.post(this.api.login,body,header)
    }
}`
