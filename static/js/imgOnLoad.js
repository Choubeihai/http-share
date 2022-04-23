function imgOnLoad() {
    let duration = window.performance.now();
    // 向id为duration的对象中插入
    document.getElementById("show-duration").innerHTML = 'time elapsed: ' + ((duration) / 1000).toFixed(10) + '秒'

}

