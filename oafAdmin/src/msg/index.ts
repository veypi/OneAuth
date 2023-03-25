/*
 * index.ts
 * Copyright (C) 2022 veypi <veypi@qq.com>
 * 2022-05-27 18:06
 * Distributed under terms of the MIT license.
 */


class Message {

    private box: HTMLDivElement
    private timeout: number
    private id: string

    constructor(id: string, timeout?: number) {
        this.timeout = timeout || 1000
        this.box = document.getElementById(id) as HTMLDivElement
        this.id = id
        if (!this.box) {
            console.error('can not found element ' + id)
            return
        }
        this.box.classList.add('v-msg-box')
        this.box.classList.add('v-msg-box-none-event')
    }
    private base(text: string, classList: string[], timeout: number) {
        let msg = document.createElement('div')
        msg.classList.add('v-msg-item')
        msg.classList.add(...classList)
        msg.innerText = text
        this.box.appendChild(msg)
        setTimeout(() => {
            msg.classList.add('v-msg-item-remove')
            msg.onanimationend = () => {
                this.box.removeChild(msg)
            }
        }, timeout)
    }
    Prompt(text: string, placeholder?: string, defaul?: string) {
        let that = this
        this.box.classList.remove('v-msg-box-none-event')
        this.box.classList.add('v-msg-box-mask')
        return new Promise(function(resolve, reject) {
            let msg = document.createElement('div')
            let title = createElement('div', ['v-msg-title'])
            let input = createElement('input', ['v-msg-input']) as HTMLInputElement
            input.placeholder = placeholder || ''

            let btn = createElement('div', [])
            let btn_ok = createElement('div', ['v-msg-ok'])
            input.value = defaul || ''
            btn_ok.innerText = "ok"
            title.innerText = text
            msg.classList.add('v-msg-prompt')
            msg.appendChild(title)
            msg.appendChild(input)
            btn.appendChild(btn_ok)
            msg.appendChild(btn)
            that.box.appendChild(msg)
            let cancel = () => {
                msg.classList.add('v-msg-item-remove')
                that.box.classList.add('v-msg-box-none-event')
                that.box.classList.remove('v-msg-box-mask')
                msg.onanimationend = () => {
                    that.box.removeChild(msg)
                }
            }
            btn_ok.onclick = () => {
                cancel()
                resolve(input.value)
            }
            that.box.onclick = (e: any) => {
                if (e.target.id === that.id) {
                    cancel()
                    reject()
                }
            }
        }).catch(e => {
            if (e) {
                console.log(e)
            }
        })
    }
    Warn(text: string) {
        this.base(text, ['v-msg-warn'], this.timeout + 1500)
    }
    Info(text: string) {
        this.base(text, ['v-msg-info'], this.timeout + 500)
    }
}

function createElement(name: string, classList: string[]) {
    let d = document.createElement(name)
    d.classList.add(...classList)
    return d
}

export { Message }

let msg: Message

function defaultMessage() {
    console.log('init message')
    if (!msg) {
        msg = new Message('v-msg')
    }
    return msg
}

export default defaultMessage()
