//liff init
const lifftest = document.getElementById('lifftest');
liff.init(
    
    data => {
        const userId = data.context.userId;
        lifftest.innerText = userId;
    },
    err => {
        console.log('error', err);
    }
);