"use strict";(()=>{function a(t,n){let e;n?e=n.querySelectorAll(t):e=document.querySelectorAll(t);let o=[];return e.forEach(r=>{o.push(r)}),o}function E(t,n){let e=a(t,n);switch(e.length){case 0:return;case 1:return e[0];default:console.warn(`found [${e.length}] elements with selector [${t}], wanted zero or one`)}}function h(t,n){let e=E(t,n);if(!e)throw new Error(`no element found for selector [${t}]`);return e}function k(t,n,e="block"){return typeof t=="string"&&(t=h(t)),t.style.display=n?e:"none",t}function se(t,...n){let e=document.createElement("li");e.innerText=t;for(let o of n){let r=document.createElement("pre");typeof o=="string"?r.innerText=o:r.innerText=JSON.stringify(o,null,2),e.appendChild(r)}return e}function $(t,...n){let e=E("#audit-log");e?e.appendChild(se(t,...n)):console.log(`### Audit ###
`+t,...n)}function q(){for(let t of a(".menu-container .final"))t.scrollIntoView({block:"center"})}var x="mode-light",H="mode-dark";function R(){for(let t of a(".mode-input"))t.onclick=()=>{switch(t.value){case"":document.body.classList.remove(x),document.body.classList.remove(H);break;case"light":document.body.classList.add(x),document.body.classList.remove(H);break;case"dark":document.body.classList.remove(x),document.body.classList.add(H);break}}}function _(t){setTimeout(()=>{t.style.opacity="0",setTimeout(()=>t.remove(),500)},5e3)}function W(t,n,e){let o=document.getElementById("flash-container");o===null&&(o=document.createElement("div"),o.id="flash-container",document.body.insertAdjacentElement("afterbegin",o));let r=document.createElement("div");r.className="flash";let i=document.createElement("input");i.type="radio",i.style.display="none",i.id="hide-flash-"+t,r.appendChild(i);let s=document.createElement("label");s.htmlFor="hide-flash-"+t;let l=document.createElement("span");l.innerHTML="\xD7",s.appendChild(l),r.appendChild(s);let c=document.createElement("div");c.className="content flash-"+n,c.innerText=e,r.appendChild(c),o.appendChild(r),_(r)}function j(){let t=document.getElementById("flash-container");if(t===null)return W;let n=t.querySelectorAll(".flash");if(n.length>0)for(let e of n)_(e);return W}function B(){for(let t of a(".link-confirm"))t.onclick=()=>{let n=t.dataset.message;return n&&n.length===0&&(n="Are you sure?"),confirm(n)}}function w(t,n){let e=(t||"").replace(/-/gu,"/").replace(/[TZ]/gu," ")+" UTC",o=new Date(e),r=(new Date().getTime()-o.getTime())/1e3,i=Math.floor(r/86400),s=o.getFullYear(),l=o.getMonth()+1,c=o.getDate();if(isNaN(i)||i<0||i>=31)return s.toString()+"-"+(l<10?"0"+l.toString():l.toString())+"-"+(c<10?"0"+c.toString():c.toString());let u,d;return i===0?r<5?(d=1,u="just now"):r<60?(d=1,u=Math.floor(r)+" seconds ago"):r<120?(d=10,u="1 minute ago"):r<3600?(d=30,u=Math.floor(r/60)+" minutes ago"):r<7200?(d=60,u="1 hour ago"):(d=60,u=Math.floor(r/3600)+" hours ago"):i===1?(d=600,u="yesterday"):i<7?(d=600,u=i+" days ago"):(d=6e3,u=Math.ceil(i/7)+" weeks ago"),n&&(n.innerText=u,setTimeout(()=>w(t,n),d*1e3)),u}function G(){return a(".reltime").forEach(t=>{w(t.dataset.time||"",t)}),w}function ce(t,n){let e=0;return(...o)=>{e!==0&&window.clearTimeout(e),e=window.setTimeout(()=>{t(null,...o)},n)}}function ae(t,n,e,o,r){var U;if(!t)return;let i=t.id+"-list",s=document.createElement("datalist"),l=document.createElement("option");l.value="",l.innerText="Loading...",s.appendChild(l),s.id=i,(U=t.parentElement)==null||U.prepend(s),t.setAttribute("autocomplete","off"),t.setAttribute("list",i);let c={},u="";function d(m){let f=n;return f.includes("?")?f+"&t=json&"+e+"="+encodeURIComponent(m):f+"?t=json&"+e+"="+encodeURIComponent(m)}function N(m){let f=c[m];!f||!f.frag||(u=m,s.replaceChildren(f.frag.cloneNode(!0)))}function re(){let m=t.value;if(m.length===0)return;let f=d(m),y=!m||!u;if(!y){let p=c[u];p&&(y=!p.data.find(T=>m===r(T)))}if(y){if(c[m]&&c[m].url===f){N(m);return}fetch(f,{credentials:"include"}).then(p=>p.json()).then(p=>{if(!p)return;let T=Array.isArray(p)?p:[p],O=document.createDocumentFragment(),A=10;for(let M=0;M<T.length&&A>0;M++){let F=r(T[M]),ie=o(T[M]);if(F){let L=document.createElement("option");L.value=F,L.innerText=ie,O.appendChild(L),A--}}c[m]={url:f,data:T,frag:O,complete:!1},N(m)})}}t.oninput=ce(re,250),console.log("managing ["+t.id+"] autocomplete")}function J(){return ae}function P(){document.addEventListener("keydown",t=>{t.key==="Escape"&&document.location.hash.startsWith("#modal-")&&(document.location.hash="")})}function I(t,n,e){return n||(n=18),e==null&&(e="icon"),`<svg class="${e}" style="width: ${n}px; height: ${n+"px"};"><use xlink:href="#svg-${t}"></use></svg>`}function le(t,n){return t.parentElement!==n.parentElement?null:t===n?0:t.compareDocumentPosition(n)&Node.DOCUMENT_POSITION_FOLLOWING?-1:1}var b;function S(t){let n=[],e=a(".item .value",t);for(let r of e)n.push(r.innerText);let o=h("input.result",t);o.value=n.join(", ")}function ue(t){var e;let n=(e=t.parentElement)==null?void 0:e.parentElement;t.remove(),n&&S(n)}function K(t){let n=h(".value",t),e=h(".editor",t);e.value=n.innerText;let o=()=>{var i;if(e.value===""){t.remove();return}n.innerText=e.value,k(n,!0),k(e,!1);let r=(i=t.parentElement)==null?void 0:i.parentElement;r&&S(r)};e.onblur=o,e.onkeydown=r=>{if(r.code==="Enter")return r.preventDefault(),o(),!1},k(n,!1),k(e,!0),e.focus()}function Q(t,n){let e=document.createElement("div");e.className="item",e.draggable=!0,e.ondragstart=s=>{var l;(l=s.dataTransfer)==null||l.setDragImage(document.createElement("div"),0,0),e.classList.add("dragging"),b=e},e.ondragover=()=>{var c;let s=le(e,b);if(!s)return;let l=s===-1?e:e.nextSibling;(c=b.parentElement)==null||c.insertBefore(b,l),S(n)},e.ondrop=s=>{s.preventDefault()},e.ondragend=s=>{e.classList.remove("dragging"),s.preventDefault()};let o=document.createElement("div");o.innerText=t,o.className="value",o.onclick=()=>{K(e)},e.appendChild(o);let r=document.createElement("input");r.className="editor",e.appendChild(r);let i=document.createElement("div");return i.innerHTML=I("times",13),i.className="close",i.onclick=()=>{ue(e)},e.appendChild(i),e}function me(t,n){let e=Q("",n);t.appendChild(e),K(e)}function Y(t){var i;let n=h("input.result",t),e=h(".tags",t),o=n.value.split(",").map(s=>s.trim()).filter(s=>s!=="");k(n,!1),e.innerHTML="";for(let s of o)e.appendChild(Q(s,t));(i=E(".add-item",t))==null||i.remove();let r=document.createElement("div");r.className="add-item",r.innerHTML=I("plus",22),r.onclick=()=>{me(e,t)},t.insertBefore(r,h(".clear",t))}function V(){for(let t of a(".tag-editor"))Y(t);return Y}var Z="--selected";function de(t){var e,o;let n=(o=(e=t.parentElement)==null?void 0:e.parentElement)==null?void 0:o.querySelector("input");if(!n)throw new Error("no associated input found");n.value="\u2205"}function C(t){t.onreset=()=>C(t);let n={},e={};for(let o of t.elements){let r=o;if(r.name.length>0)if(r.name.endsWith(Z))e[r.name]=r;else{(r.type!=="radio"||r.checked)&&(n[r.name]=r.value);let i=()=>{let s=e[r.name+Z];s&&(s.checked=n[r.name]!==r.value)};r.onchange=i,r.onkeyup=i}}}function X(){for(let t of a("form.editor"))C(t);return[de,C]}var fe=[];function z(t,n,e){let o=t.querySelectorAll(n);if(o.length===0)throw new Error("empty query selector ["+n+"]");o.forEach(r=>e(r))}function v(t,n,e){z(t,n,o=>{o.style.backgroundColor=e})}function g(t,n,e){z(t,n,o=>{o.style.color=e})}function pe(t,n,e){let o=document.querySelector("#mockup-"+t);if(!o){console.error("can't find mockup for mode ["+t+"]");return}switch(n){case"color-foreground":g(o,".mock-main",e);break;case"color-background":v(o,".mock-main",e);break;case"color-foreground-muted":g(o,".mock-main .mock-muted",e);break;case"color-background-muted":v(o,".mock-main .mock-muted",e);break;case"color-link-foreground":g(o,".mock-main .mock-link",e);break;case"color-link-visited-foreground":g(o,".mock-main .mock-link-visited",e);break;case"color-nav-foreground":g(o,".mock-nav",e),g(o,".mock-nav .mock-link",e);break;case"color-nav-background":v(o,".mock-nav",e);break;case"color-menu-foreground":g(o,".mock-menu",e),g(o,".mock-menu .mock-link",e);break;case"color-menu-background":v(o,".mock-menu",e);break;case"color-menu-selected-foreground":g(o,".mock-menu .mock-link-selected",e);break;case"color-menu-selected-background":v(o,".mock-menu .mock-link-selected",e);break;default:console.error("invalid key ["+n+"]")}}function ee(){for(let t of a(".color-var")){let n=t.dataset.var,e=t.dataset.mode;fe.push(n),!(!n||n.length===0)&&(t.oninput=()=>{pe(e,n,t.value)})}}var te=!1;function ne(t){if(t||(t="/connect"),t.indexOf("ws")===0)return t;let n=document.location,e="ws";return n.protocol==="https:"&&(e="wss"),t.indexOf("/")!==0&&(t="/"+t),e+`://${n.host}${t}`}var D=class{constructor(n,e,o,r,i){this.debug=n,this.open=e,this.recv=o,this.err=r,this.url=ne(i),this.connected=!1,this.pauseSeconds=1,this.pendingMessages=[],this.connect()}connect(){window.onbeforeunload=()=>{te=!0},this.connectTime=Date.now(),this.sock=new WebSocket(ne(this.url));let n=this;this.sock.onopen=()=>{n.connected=!0,n.pendingMessages.forEach(n.send),n.pendingMessages=[],n.debug&&console.log("WebSocket connected"),n.open()},this.sock.onmessage=e=>{let o=JSON.parse(e.data);n.debug&&console.debug("[socket]: receive",o),n.recv(o)},this.sock.onerror=e=>()=>{n.err("socket",e.type)},this.sock.onclose=()=>{if(te)return;n.connected=!1;let e=n.connectTime?Date.now()-n.connectTime:0;e>0&&e<2e3?(n.pauseSeconds=n.pauseSeconds*2,n.debug&&console.debug(`socket closed immediately, reconnecting in ${n.pauseSeconds} seconds`),setTimeout(()=>{n.connect()},n.pauseSeconds*1e3)):(console.debug("socket closed after ["+e+"ms]"),setTimeout(()=>{n.connect()},n.pauseSeconds*500))}}disconnect(){}send(n){if(this.debug&&console.debug("out",n),!this.sock)throw new Error("socket not initialized");if(this.connected){let e=JSON.stringify(n,null,2);this.sock.send(e)}else this.pendingMessages.push(n)}};function oe(){return D}function ge(){let[t,n]=X();window.loadtoad={relativeTime:G(),autocomplete:J(),setSiblingToNull:t,initForm:n,flash:j(),tags:V(),Socket:oe()},q(),R(),B(),P(),ee(),window.audit=$}document.addEventListener("DOMContentLoaded",ge);})();
//# sourceMappingURL=client.js.map
