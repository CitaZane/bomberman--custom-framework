import createRouter from "../framework/router";
import createStore from "../framework/store";
import routes from "./router";
import storeObj from "./store/index";

const store = createStore(storeObj);
const router = createRouter(routes);

document.addEventListener("keyup", (e) =>{
        store.dispatch('registerKeyUp', e.key)
    })
document.addEventListener("keydown", (e) =>{
        store.dispatch('registerKeyDown', e.key)
    })
export { store, router };