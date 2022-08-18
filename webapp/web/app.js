function getTelegramInitData() {
    window.Telegram.WebApp.ready();
    return window.Telegram.WebApp.initData;
}

function getTelegramThemeParams() {
    return JSON.stringify(window.Telegram.WebApp.themeParams);
}

function getTelegramColorScheme() {
    return window.Telegram.WebApp.colorScheme;
}

window.logger = (flutter_value) => {
   console.log({ js_context: this, flutter_value });
}
