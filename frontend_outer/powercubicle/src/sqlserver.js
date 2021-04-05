import axios from 'axios'
import { setCookie, getCookie } from './localstorage'

axios.defaults.timeout = 5000
axios.defaults.baseURL = ''

const BACKENDHOST = 'http://139.198.15.216:12077'
const BACKENDDBHOST = 'http://139.198.15.216:5001'
const TOKEN = 'PCSessionToken'
const USERNAME = 'PCUserName'
const USERSEAT = 'PCUserSeat'
// const querystring = require('querystring')
const SUCCESS = 'SUCCESS'
const ERROR = 'ERROR'
const NAMEDUP = 'NAMEDUP'
const NAMEERROR = 'NAMEERROR'
const PWDERROR = 'PWDERROR'
// const RESERVED = 'RESERVED'
// const FORMATERROR = 'FORMATERROR'

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
      resolve(response.data)
    }).catch(err => {
      reject(err)
    })
  })
}

export function getCurrentSeat () {
  return new Promise((resolve, reject) => {
    axios.get(BACKENDHOST + '/powercubicle/v1/user/seat',
      {
        headers: {
          'Authorization': 'Bearer ' + getCookie(TOKEN)
        }
      }
    ).then(response => {
      var seatid = response.data.Seat
      if (seatid !== '') {
        setCookie(USERSEAT, seatid, 10 * 60 * 60 * 1000)
        resolve(seatid)
      } else {
        setCookie(USERSEAT, '', 10 * 60 * 60 * 1000)
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
    axios.post(BACKENDHOST + '/powercubicle/v1/user/login', JSON.stringify(parseData),
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
    axios.post(BACKENDHOST + '/powercubicle/v1/user/register', JSON.stringify(parseData),
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

export function usrPwdChange (usrPwd) {
  let parseData = {
    'password': usrPwd
  }
  return new Promise((resolve, reject) => {
    axios.post(BACKENDDBHOST + '/powercubicle/v1/db/user/password', parseData,
      {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ' + getCookie(TOKEN)
        }
      }
    ).then(response => {
      console.log(response)
      if (response.status === 200) {
        resolve(SUCCESS)
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

export function seatRelease () {
  return new Promise((resolve, reject) => {
    axios.post(BACKENDHOST + '/powercubicle/v1/seat/release', null,
      {
        headers: {
          'Authorization': 'Bearer ' + getCookie(TOKEN)
        }
      }
    ).then(response => {
      console.log(response)
      setCookie(USERSEAT, '', 1000)
      resolve(SUCCESS)
    }, err => {
      console.log(err.message)
      reject(err)
    })
  })
}

export function seatRegis (encryCode) {
  let parseData = {
    'encrypted_qrcode': encryCode
  }
  return new Promise((resolve, reject) => {
    axios.post(BACKENDHOST + '/powercubicle/v1/seat/register', JSON.stringify(parseData),
      {
        headers: {
          'Authorization': 'Bearer ' + getCookie(TOKEN),
          'Content-Type': 'text/plain'
        }
      }
    ).then(response => {
      console.log(response)
      if (response.data.seat_number !== '') {
        setCookie(USERSEAT, response.data.seat_number, 10 * 60 * 60 * 1000)
        resolve(response.data.seat_number)
      } else {
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
