import axios from 'axios'
import { setCookie, getCookie } from './localstorage'

axios.defaults.timeout = 5000
axios.defaults.baseURL = ''

const BACKENDHOST = 'http://192.168.31.108:12077'
const TOKEN = 'PCSessionToken'
const USERNAME = 'PCUserName'
// const querystring = require('querystring')
const SUCCESS = 'SUCCESS'
const ERROR = 'ERROR'
const NAMEDUP = 'NAMEDUP'
const NAMEERROR = 'NAMEERROR'
const PWDERROR = 'PWDERROR'
// const RESERVED = 'RESERVED'
// const FORMATERROR = 'FORMATERROR'
// http request 拦截器
axios.interceptors.request.use(
  config => {
    // const token = getCookie('名称');注意使用的时候需要引入cookie方法，推荐js-cookie
    // config.data = JSON.stringify(config.data)
    // config.data = JSON.stringify(config.data)
    // config.headers = {
    // 'Content-Type': 'application/x-www-form-urlencoded'
    // 'Content-Type': 'application/json;charset=UTF-8'
    // }
    // if(token){
    //   config.params = {'token':token}
    // }
    return config
  }, err => {
    return Promise.reject(err)
  }
)

/**
 * 封装get方法
 * @param url
 * @param data
 * @returns {Promise}
 */
export function get (url, params = {}) {
  return new Promise((resolve, reject) => {
    axios.get(url, {
      params: params
    }).then(response => {
      resolve(response.data)
    }).catch(err => {
      reject(err)
    })
  })
}

export function getSeatNum () {
  return new Promise((resolve, reject) => {
    axios.get(BACKENDHOST + '/powercubicle/v1/seat').then(response => {
      resolve(response.data.Data.Seat)
    }).catch(err => {
      reject(err)
    })
  })
}

export function getSeatSitu (loc) {
  return new Promise(function (resolve, reject) {
    axios.get(BACKENDHOST + '/powercubicle/v1/seat?region=' + loc).then(response => {
      resolve(response.data.Data.Seat)
    }).catch(err => {
      reject(err)
    })
  })
}

export function getCurrentSeat () {
  return new Promise((resolve, reject) => {
    axios.get(BACKENDHOST + '/Reservation/user/seat', {
      header: {
        'Authorization': 'Bearer ' + getCookie(TOKEN),
        'content-type': 'application/json'
      }
    }).then(response => {
      resolve(response.data) // TODO
    }).catch(err => {
      reject(err)
    })
  })
}

/**
 * 封装post请求
 * @param url
 * @param data
 * @returns {Promise}
 */
export function post (url, data = {}) {
  return new Promise((resolve, reject) => {
    axios.post(url, data)
      .then(response => {
        resolve(response.data)
      }, err => {
        reject(err)
      })
  })
}

export function usrLogin (usrName, usrPwd) {
  let parseData = {
    'name': usrName,
    'password': usrPwd
  }
  return new Promise((resolve, reject) => {
    console.log(usrName, usrPwd)
    axios.post(BACKENDHOST + '/powercubicle/v1/user/login', parseData,
      {
        headers: {
          'Content-Type': 'text/plain'
        }
      }
    ).then(response => {
      console.log(response)
      var token = response.data.Session_key
      if (typeof (token) !== 'undefined') {
        setCookie(TOKEN, token, 30 * 24 * 60 * 60 * 1000)
        setCookie(USERNAME, usrName, 30 * 24 * 60 * 60 * 1000)
        this.$router.push({
          path: '/'
        })
        resolve(SUCCESS)
      } else if (response.data.Data.message.indexOf('username') >= 0) {
        resolve(NAMEERROR)
      } else if (response.data.Data.message.indexOf('password') >= 0) {
        resolve(PWDERROR)
      } else {
        console.log(ERROR)
        resolve(ERROR)
      }
    }, err => {
      console.log(err)
      reject(err)
    })
  })
}

export function usrRegis (usrName, usrPwd) {
  let parseData = {
    'name': usrName,
    'password': usrPwd
  }
  return new Promise((resolve, reject) => {
    console.log(usrName, usrPwd)
    axios.post(BACKENDHOST + '/powercubicle/v1/user/register', parseData,
      {
        headers: {
          'Content-Type': 'text/plain'
        }
      }
    ).then(response => {
      console.log(response)
      var token = response.data.Session_key
      if (typeof (token) !== 'undefined') {
        setCookie(TOKEN, token, 30 * 24 * 60 * 60 * 1000)
        setCookie(USERNAME, usrName, 30 * 24 * 60 * 60 * 1000)
        this.$router.push({
          path: '/'
        })
        resolve(SUCCESS)
      } else if (response.data.Data.message.indexOf('used') >= 0) {
        resolve(NAMEDUP)
      } else {
        console.log(ERROR)
        resolve(ERROR)
      }
    }, err => {
      console.log(err)
      reject(err)
    })
  })
}

export function userSeat () {
  return new Promise((resolve, reject) => {
    axios.get(BACKENDHOST + '/powercubicle/v1/user/seat',
      {
        headers: {
          'Authorization': 'Bearer ' + getCookie(TOKEN),
          'Content-Type': 'application/json'
        }
      }
    ).then(response => {
      console.log(response)
      var token = response.data.Session_key
      if (typeof (token) !== 'undefined') {
        setCookie(TOKEN, token, 30 * 24 * 60 * 60 * 1000)
        setCookie(USERNAME, usrName, 30 * 24 * 60 * 60 * 1000)
        this.$router.push({
          path: '/'
        })
        resolve(SUCCESS)
      } else if (response.data.Data.message.indexOf('used') >= 0) {
        resolve(NAMEDUP)
      } else {
        console.log(ERROR)
        resolve(ERROR)
      }
    }, err => {
      console.log(err)
      reject(err)
    })
  })
}
/**
 * 封装patch请求
 * @param url
 * @param data
 * @returns {Promise}
 */
export function patch (url, data = {}) {
  return new Promise((resolve, reject) => {
    axios.patch(url, data).then(response => {
      resolve(response.data)
    }, err => {
      reject(err)
    })
  })
}

/**
 * 封装put请求
 * @param url
 * @param data
 * @returns {Promise}
 */
export function put (url, data = {}) {
  return new Promise((resolve, reject) => {
    axios.put(url, data).then(response => {
      resolve(response.data)
    }, err => {
      reject(err)
    })
  })
}
