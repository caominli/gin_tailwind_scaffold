function theme () {
    var theme1 = localStorage.getItem('theme'); //读取本地主题
    if (theme1 != "dark") {
    document.documentElement.classList.remove('dark');
    console.log("当前为亮光")
    return false;
    } else {
    document.documentElement.classList.add('dark');
    console.log("当前为黑暗")
    return true;
    }
}
