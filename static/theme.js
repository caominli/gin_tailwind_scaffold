//读取主题，并添加样式，返回是否暗黑模式
function theme () {
    var theme1 = localStorage.getItem('theme'); //读取本地主题
    if (theme1 != 'dark') {
    document.documentElement.classList.remove('dark');
    return false;
    } else {
    document.documentElement.classList.add('dark');
    return true;
    }
}

//消息闪现
function Message (app,msg,type) {
    var svg;
    switch (type) {
        case "1":
            svg = `
                <svg viewBox="0 0 20 20" fill="currentColor" class="h-6 w-6 text-rose-600">
                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.28 7.22a.75.75 0 00-1.06 1.06L8.94 10l-1.72 1.72a.75.75 0 101.06 1.06L10 11.06l1.72 1.72a.75.75 0 101.06-1.06L11.06 10l1.72-1.72a.75.75 0 00-1.06-1.06L10 8.94 8.28 7.22z" clip-rule="evenodd" />
                </svg>         
            `;
            break;

        case "2":
            svg = `
                <svg viewBox="0 0 20 20" fill="currentColor" class="h-6 w-6 text-yellow-500">
                    <path fill-rule="evenodd" d="M8.485 2.495c.673-1.167 2.357-1.167 3.03 0l6.28 10.875c.673 1.167-.17 2.625-1.516 2.625H3.72c-1.347 0-2.189-1.458-1.515-2.625L8.485 2.495zM10 5a.75.75 0 01.75.75v3.5a.75.75 0 01-1.5 0v-3.5A.75.75 0 0110 5zm0 9a1 1 0 100-2 1 1 0 000 2z" clip-rule="evenodd" />
                </svg>           
            `;
            break;
        default:
            svg = `
                <svg class="h-6 w-6 text-green-400" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>         
            `;
            break;
    };

   
    var html = `
        <div v-if="showMsg" aria-live="assertive" class="pointer-events-none fixed inset-0 flex px-4 py-6 items-start sm:p-6 z-10">
            <div class="flex w-full flex-col items-center space-y-4 sm:items-end">
            <div class="pointer-events-auto w-full max-w-sm overflow-hidden rounded-lg bg-white shadow-lg ring-1 ring-black ring-opacity-5">
                <div class="p-4">
                <div class="flex items-start">
                    <div class="flex-shrink-0">
                    ${svg}
                    </div>
                    <div class="ml-3 w-0 flex-1 pt-0.5">
                    <p class="text-sm font-medium text-gray-900">${msg}</p>
                    </div>
                </div>
                </div>
            </div>
            </div>
        </div>
    `;
    app.msg = html;
    //定时关闭
    setTimeout(() => {
        app.msg = '';
      }, 3500);
}

//消息闪现调用方法：Message(this,"这是一个测试",2); 第一个参数要在vue中传入自身，第二个参数 是消息内容，第三个参数：1错误，2警告，3完成默认
