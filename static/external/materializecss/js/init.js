$(document).ready(function() {
    $('.parallax').parallax();
    $('.modal-trigger').leanModal();

    $('.datepicker').pickadate({
        min:"today",
        selectMonths: true, // Creates a dropdown to control month
        selectYears: 2 // Creates a dropdown of 15 years to control year
    });
});
