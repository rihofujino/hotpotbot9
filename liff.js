window.onload = function (e) {
    // init で初期化。基本情報を取得。
    // https://developers.line.me/ja/reference/liff/#initialize-liff-app
    liff.init(function (data) {
        initializeApp(data);
    });
};

