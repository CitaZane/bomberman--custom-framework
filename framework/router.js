import diff from "./vDom/diff";
import render, { refs, onMountedStack } from "./vDom/render";
import mount from "./vDom/mount";
import { watchEffect } from "./reactive";

class Router {
  #routerMap = []; //holds all registered routes
  #defaultRoute = {}; //fallback route
  currentRoute = {}; //current route -> to be accesed in component if needed

  $rootElem;
  $app;
  vApp;

  constructor(routes) {
    // import all defined routes add to router
    this.registerRoutes(routes);
    // set default route as the first in list
    this.#defaultRoute = this.#routerMap[0];
    // add listener for router
    window.onhashchange = this.updateView.bind(this);
  }

  updateDom() {
    if (!this.vApp) {
      this.updateView();
      return;
    }
    const view = this.currentRoute.component();

    const vNewApp = view.template;
    const patch = diff(this.vApp, vNewApp);
    this.$rootElem = patch(this.$rootElem);
    // trigger noMounted hook
    view?.onMounted?.call(null, refs);

    onMountedStack.forEach((fn) => fn.call(null, refs));
    onMountedStack.length = 0;

    this.vApp = vNewApp;
  }

  updateView() {
    // get route name from address bar
    var routeName = location.hash.replace("#", "");
    let _res = this.matchRoutes(routeName);
    if (!this.vApp) {
      const view = this.currentRoute.component();
      view?.created?.getPlayerCount();
      this.vApp = view.template;
      this.$app = render(this.vApp);
      this.$rootElem = mount(this.$app, document.getElementById("root"));
      // trigger noMounted hook
      view?.onMounted?.call(null, refs);
      onMountedStack.forEach((fn) => fn());
      onMountedStack.length = 0; //clear
    } else {
      this.updateDom();
    }
  }

  // based on provided url path -> find match return it || return default
  matchRoutes(routeName) {
    const matches = [];
    // check each registered route against provided
    this.#routerMap.forEach((r) => {
      let match = routeName.match(r.path);

      if (match) {
        const params = match.slice(2); //access provided params
        // assign params to keys
        var finalParams = {};
        r.keys.forEach((k, i) => {
          finalParams[k] = params[i];
        });
        // set match as current route
        this.currentRoute = { ...r, params: finalParams };
        matches.push({ ...r, params: finalParams });
        return;
      }
    });
    // in case of no match found set default view as current
    if (!matches[0]) {
      this.currentRoute = this.#defaultRoute;
      // replace url with appropriate path name
      window.history.pushState("", "", this.#defaultRoute.name);
    }
  }

  registerRoutes(routes) {
    routes.forEach((route) => {
      let newRoute = new Route(route);
      this.#routerMap.push(newRoute);
    });
  }
}

// route consists of -> name -> same as user defined path
//                   -> path -> regex for matching path
//                   -> keys -> [] of all prop names
//                   -> props -> {} of key name and prop from url
//                   -> compnent -> to be rendered
class Route {
  constructor(route) {
    var pathCompiled = this.compilePath(route.path);

    this.name = route.path;
    this.path = pathCompiled.regex; //holds regex expression for path
    this.keys = pathCompiled.keys;
    this.component = route.component;
  }

  compilePath(path) {
    const keys = [];

    path = path.replace(/:(\w+)/g, (_, key) => {
      keys.push(key);
      return "([^\\/]+)";
    });

    const source = `^(${path})$`;
    const regex = new RegExp(source, "i");

    return { regex, keys };
  }
}

export default function createRouter(routes) {
  let router = new Router(routes);

  // start listening for changes
  watchEffect(() => {
    router.updateDom();
  });

  return router;
}
