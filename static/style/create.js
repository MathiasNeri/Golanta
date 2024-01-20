
function updateSliders(changedSlider) {
    var hpSlider = document.getElementById('hp');
    var resSlider = document.getElementById('res');
    var hpValue = document.getElementById('hp-value');
    var resValue = document.getElementById('res-value');

    var hp = parseInt(hpSlider.value);
    var res = parseInt(resSlider.value);

    // Mettez à jour les valeurs max des curseurs en fonction de la somme totale
    var totalPoints = 10;
    if (changedSlider.id === 'hp') {
        resSlider.max = totalPoints - hp;
    } else if (changedSlider.id === 'res') {
        hpSlider.max = totalPoints - res;
    }

    // Mise à jour des valeurs affichées
    hpValue.textContent = hp;
    resValue.textContent = res;
}

// Initialisez les curseurs correctement lors du chargement de la page
document.addEventListener('DOMContentLoaded', function () {
    updateSliders(document.getElementById('hp'));
});
