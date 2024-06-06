/*
 * state.ts
 * Copyright (C) 2024 veypi <i@veypi.com>
 * 2024-06-06 14:44
 * Distributed under terms of the MIT license.
 */


import { type Auths } from '@/composables'
interface localuser {
  id: string
  local: modelsUser
  auth: Auths
  ready: boolean
}


export const useLocalUser = () => useState<localuser>('user', () => ({
  id: '',
  local: {} as modelsUser,
  auth: {} as Auths,
  ready: false
}))


interface Vapp {
  id: string;
  layout: {
    theme: string;
    // layout px
    header_height: number;
    footer_height: number;
    menu_width: number;
  }
}


export const useVapp = () => useState<Vapp>('vapp', () => ({
  id: '',
  layout: {
    theme: '',
    header_height: 80,
    footer_height: 16,
    menu_width: 40,
  }
}))

