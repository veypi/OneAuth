function padLeftZero(str: string): string {
  return ('00' + str).substr(str.length)
}

const util = {
  title: function (title: string) {
    window.document.title = title ? title + ' - oa' : 'veypi project'
  },
  getCookie(name: string) {
    const reg = new RegExp('(^| )' + name + '=([^;]*)(;|$)')
    const arr = document.cookie.match(reg)
    if (arr) {
      return unescape(arr[2])
    } else return null
  },
  delCookie(name: string) {
    const exp = new Date()
    exp.setTime(exp.getTime() - 1)
    const cval = this.getCookie(name)
    if (cval !== null) {
      document.cookie = name + '=' + cval + ';expires=' + exp.toLocaleString()
    }
  },
  setCookie(name: string, value: string, time: number) {
    const exp = new Date()
    exp.setTime(exp.getTime() + time)
    document.cookie =
      name + '=' + escape(value) + ';expires=' + exp.toLocaleString()
  },
  checkLogin() {
    // return parseInt(this.getCookie('stat')) === 1
    return Boolean(localStorage.auth_token)
  },

  formatDate(date: Date, fmt: string) {
    if (/(y+)/.test(fmt)) {
      fmt = fmt.replace(
        RegExp.$1,
        (date.getFullYear() + '').substr(4 - RegExp.$1.length)
      )
    }
    const o = {
      'M+': date.getMonth() + 1,
      'd+': date.getDate(),
      'h+': date.getHours(),
      'm+': date.getMinutes(),
      's+': date.getSeconds()
    }
    for (const k in o) {
      if (new RegExp(`(${k})`).test(fmt)) {
        // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
        // @ts-ignore
        const str = o[k] + ''
        fmt = fmt.replace(
          RegExp.$1,
          RegExp.$1.length === 1 ? str : padLeftZero(str)
        )
      }
    }
    return fmt
  }
}

export default util
