import jsx from "../framework/vDom/jsx";

export function Overlay() {
    var timeleft = 10;
    var intervalId = setInterval(function () {
        if (timeleft <= 0) {
            clearInterval(intervalId);
            document.getElementById("overlay").style.visibility="hidden";
        }
        document.getElementById("overlay-counter").textContent = timeleft;
        timeleft -= 1;
    }, 1000);
    return {
        template: (
            <div id="overlay">
                <div id="overlay-counter"></div>
            </div>
        ),
    };
}