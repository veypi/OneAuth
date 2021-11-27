import {darkTheme} from 'naive-ui/lib/themes'
import {BuiltInGlobalTheme} from 'naive-ui/lib/themes/interface'
import {lightTheme} from 'naive-ui/lib/themes/light'
import {ref} from 'vue'
import {useOsTheme, GlobalThemeOverrides} from 'naive-ui'

interface builtIn extends BuiltInGlobalTheme {
    overrides: GlobalThemeOverrides
    me: {
        lightBox: string,
        lightBoxShadow: string
    }
}

let light = lightTheme as builtIn
let dark = darkTheme as builtIn
let intputNone = {
    color: 'url(0) no-repeat',
    colorFocus: 'url(0) no-repeat',
    colorFocusWarning: 'url(0) no-repeat',
    colorFocusError: 'url(0) no-repeat',
}
light.overrides = {
    Input: Object.assign({}, intputNone),
}
dark.overrides = {
    Input: Object.assign({
        border: '1px solid #aaa',
    }, intputNone),
}
light.common.cardColor = '#f4f4f4'
light.common.bodyColor = '#eee'
dark.common.bodyColor = '#2e2e2e'
light.me = {
    lightBox: '#f4f4f4',
    lightBoxShadow: '18px 18px 36px #c6c6c6, -18px -18px 36px #fff',
}

dark.me = {
    lightBox: '#2e2e2e',
    lightBoxShadow: '21px 21px 42px #272727, -21px -21px 42px #353535',
}
export const OsThemeRef = useOsTheme()

let theme = 'light'

export let Theme = ref(light)

export let IsDark = ref(false)

function change(t: string) {
    if (t === 'dark') {
        theme = 'dark'
        Theme.value = dark
    } else {
        theme = 'light'
        Theme.value = light
    }
    IsDark.value = theme === 'dark'
}

export function ChangeTheme() {
    if (IsDark.value) {
        change('light')
    } else {
        change('dark')
    }
}

if (OsThemeRef.value === 'dark') {
    change('dark')
}
