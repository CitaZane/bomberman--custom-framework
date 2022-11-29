/* @jsx jsx */

import jsx from "../framework/vDom/jsx";

export function Overlay({ content }) {
  return {
    template: (
      <div id="overlay">
        <div class="overlay-content" id="winner">
          {content}
        </div>
      </div>
    ),
  };
}
// var timeleft = 10;
// var intervalId = setInterval(function () {
//     if (timeleft <= 0) {
//         clearInterval(intervalId);
//         document.getElementById("overlay").style.visibility="hidden";
//     }
//     document.getElementById("overlay-counter").textContent = timeleft;
//     timeleft -= 1;
// }, 1000);
