// simple as that

const backButton = document.body.querySelector('#backButton');
backButton.addEventListener('click', event => {
    history.back();
});