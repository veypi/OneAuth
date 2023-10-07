/*
 * fs.ts
 * Copyright (C) 2023 veypi <i@veypi.com>
 * 2023-10-08 01:55
 * Distributed under terms of the MIT license.
 */


import axios from "axios";
import { Base64 } from 'js-base64'
import util from "./util";
import { createClient, WebDAVClient } from 'webdav'

export interface fileProps {
  filename: string,
  basename: string,
  lastmod: string,
  size: number,
  type: "directory" | "file",
  etag: string
}

let cfg = {
  token: '',
  host: '',
  dav: {} as WebDAVClient,
}

const setCfg = (token: string) => {
  cfg.token = token
  cfg.dav = createClient('/file/',
    { headers: { auth_token: cfg.token } })
}

const rename = (o: string, n?: string) => {
  let ext = '.' + o.split('.').pop()?.toLowerCase()
  if (n) {
    return n + ext
  }
  let d = new Date().getTime()
  return d + Base64.encode(o) + ext
}


const get = (url: string): Promise<string> => {
  return fetch(url, { headers: { auth_token: util.getToken() } }).then((response) => response.text())
}

const upload = (f: FileList | File[], renames?: string[]) => {
  return new Promise<string[]>((resolve, reject) => {
    var data = new FormData();
    for (let i = 0; i < f.length; i++) {
      let nf = renames ? new File([f[i]], rename(f[i].name, renames[i]), { type: f[i].type }) : f[i]
      data.append('files', nf, nf.name)
    }
    axios.post("/api/upload/", data, {
      headers: {
        "Content-Type": 'multipart/form-data',
        'auth_token': cfg.token,
      }
    }).then(e => {
      resolve(e.data)
    }).catch(reject)
  })
}


const dav = () => {
  return {
    stat: cfg.dav.stat,
    dir: cfg.dav.getDirectoryContents,
    upload: (dir: string, name: string, file: any) => {
      return new Promise((resolve, reject) => {
        let reader = new FileReader()
        reader.onload = function(event) {
          var res = event.target?.result
          // let data = new Blob([res])
          cfg.dav.putFileContents(name, res).then(e => {
            resolve(e)
          }).catch(reject)
        }
        reader.readAsArrayBuffer(file)
      });
    }
  }
}

export default {
  setCfg,
  get,
  upload,
  dav,
}
