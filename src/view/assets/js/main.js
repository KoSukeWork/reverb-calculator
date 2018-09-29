$(document).ready(function () {
    $("form .reset_button").click(function () {
        $("form input[type=text]").val("");
        $("form").submit();
    });
})
