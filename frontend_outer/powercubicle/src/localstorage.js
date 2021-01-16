export function setCookie (key, value, expiredays) {
  var exdate = new Date()
  exdate.setDate(exdate.getDate() + expiredays)
  document.cookie = key + '=' + escape(value) + ((expiredays == null) ? '' : ';expires=' + exdate.toGMTString())
}

export function getCookie (key) {
  if (document.cookie.length > 0) {
    var cStart = document.cookie.indexOf(key + '=')
    if (cStart !== -1) {
      cStart = cStart + key.length + 1
      var cEnd = document.cookie.indexOf(';', cStart)
      if (cEnd === -1) {
        cEnd = document.cookie.length
      }
      return unescape(document.cookie.substring(cStart, cEnd))
    }
  }
  return ''
}
