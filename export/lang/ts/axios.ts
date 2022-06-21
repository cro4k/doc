import axios from "axios"


interface Login {
    username:string,
    password:string,
    remember?:boolean,
}

interface Info {
    id:string,
    data?:any,
}


export default {
    api: {
        login:"",
    },
    async post(api:string,body?:any, header?:any) {
        return axios.post(api,body, {headers:header})
    },
    async get(api:string,body?:any, header?:any)  {
        return axios.get(api,body,{headers:header})
    },
    async login(body:Login,header?:any) {
        return axios.post(this.api.login,body, {headers:header})
    }
}