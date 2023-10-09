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
  app_id: '',
  dav: {} as WebDAVClient,
  app_dav: {} as WebDAVClient,
}

const setCfg = (options: any) => {
  cfg.token = options.token
  cfg.app_id = options.app_id
  cfg.app_dav = createClient('/file/',
    { headers: { auth_token: cfg.token, app_id: options.app_id } })
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

// rename 可以保持url不变
const upload = (f: FileList | File[], dir?: string, renames?: string[]) => {
  return new Promise<string[]>((resolve, reject) => {
    var data = new FormData();
    for (let i = 0; i < f.length; i++) {
      let nf = new File([f[i]], rename(f[i].name, renames && renames[i] ? renames[i] : undefined), { type: f[i].type })
      data.append('files', nf, nf.name)
    }
    axios.post("/api/upload/" + (dir || ''), data, {
      headers: {
        "Content-Type": 'multipart/form-data',
        'auth_token': cfg.token,
      }
    }).then(e => {
      resolve(e.data)
    }).catch(reject)
  })
}

const get_dav = (client: WebDAVClient) => {
  return {
    client: client,
    stat: client.stat,
    dir: client.getDirectoryContents,
    upload: (dir: string, name: string, file: any) => {
      return new Promise((resolve, reject) => {
        let temp = () => {
          let reader = new FileReader()
          reader.onload = function(event) {
            var res = event.target?.result
            // let data = new Blob([res])
            client.putFileContents(dir + name, res).then(e => {
              if (e) {
                resolve('/file' + dir + name)
              }
            }).catch(reject)
          }
          reader.readAsArrayBuffer(file)
        }
        client.stat(dir).then(e => {
          temp()
        }).catch(e => {
          client.createDirectory(dir, { recursive: true }).then(e => {
            temp()
          }).catch(e => {
            console.log(e)
          })
        })
      });
    }
  }
}


const dav = () => {
  return get_dav(cfg.dav)
}
const appdav = () => {
  return get_dav(cfg.app_dav)
}

export default {
  setCfg,
  get,
  upload,
  dav,
  appdav,
  rename,
}
