"use strict";(()=>{function d(e,n){let t;n?t=n.querySelectorAll(e):t=document.querySelectorAll(e);let o=[];return t.forEach(r=>{o.push(r)}),o}function y(e,n){let t=d(e,n);switch(t.length){case 0:return;case 1:return t[0];default:console.warn(`found [${t.length}] elements with selector [${e}], wanted zero or one`)}}function T(e,n){let t=y(e,n);if(!t)throw new Error(`no element found for selector [${e}]`);return t}function v(e,n,t="block"){return typeof e=="string"&&(e=T(e)),e.style.display=n?t:"none",e}function ce(e,...n){let t=document.createElement("li");t.innerText=e;for(let o of n){let r=document.createElement("pre");typeof o=="string"?r.innerText=o:r.innerText=JSON.stringify(o,null,2),t.appendChild(r)}return t}function B(e,...n){let t=y("#audit-log");t?t.appendChild(ce(e,...n)):console.log(`### Audit ###
`+e,...n)}function q(){for(let e of d(".menu-container .final"))e.scrollIntoView({block:"center"})}var I="mode-light",C="mode-dark";function R(){for(let e of d(".mode-input"))e.onclick=()=>{switch(e.value){case"":document.body.classList.remove(I),document.body.classList.remove(C);break;case"light":document.body.classList.add(I),document.body.classList.remove(C);break;case"dark":document.body.classList.remove(I),document.body.classList.add(C);break}}}function W(e){setTimeout(()=>{e.style.opacity="0",setTimeout(()=>e.remove(),500)},5e3)}function $(e,n,t){let o=document.getElementById("flash-container");o===null&&(o=document.createElement("div"),o.id="flash-container",document.body.insertAdjacentElement("afterbegin",o));let r=document.createElement("div");r.className="flash";let s=document.createElement("input");s.type="radio",s.style.display="none",s.id="hide-flash-"+e,r.appendChild(s);let i=document.createElement("label");i.htmlFor="hide-flash-"+e;let a=document.createElement("span");a.innerHTML="\xD7",i.appendChild(a),r.appendChild(i);let f=document.createElement("div");f.className="content flash-"+n,f.innerText=t,r.appendChild(f),o.appendChild(r),W(r)}function _(){let e=document.getElementById("flash-container");if(e===null)return $;let n=e.querySelectorAll(".flash");if(n.length>0)for(let t of n)W(t);return $}function j(){for(let e of d(".link-confirm"))e.onclick=()=>{let n=e.dataset.message;return n&&n.length===0&&(n="Are you sure?"),confirm(n)}}function G(e){let n=e.dataset.timestamp;if(!n){console.log("invalid timestamp ["+n+"]",e);return}let t=new Date(n),o=document.createElement("span");o.title=t.toISOString(),e.classList.contains("millis")?o.innerText=t.toISOString():o.innerText=t.toLocaleString(),e.replaceChildren(o)}function S(e){let t=(e.dataset.timestamp||"").replace(/-/gu,"/").replace(/[TZ]/gu," ")+" UTC",o=new Date(t),r=(new Date().getTime()-o.getTime())/1e3,s=Math.floor(r/86400),i=o.getFullYear(),a=o.getMonth()+1,f=o.getDate();if(isNaN(s)||s<0||s>=31)return i.toString()+"-"+(a<10?"0"+a.toString():a.toString())+"-"+(f<10?"0"+f.toString():f.toString());let m,c;return s===0?r<5?(c=1,m="just now"):r<60?(c=1,m=Math.floor(r)+" seconds ago"):r<120?(c=10,m="1 minute ago"):r<3600?(c=30,m=Math.floor(r/60)+" minutes ago"):r<7200?(c=60,m="1 hour ago"):(c=60,m=Math.floor(r/3600)+" hours ago"):s===1?(c=600,m="yesterday"):s<7?(c=600,m=s+" days ago"):(c=6e3,m=Math.ceil(s/7)+" weeks ago"),e&&(e.innerText=m,setTimeout(()=>S(e),c*1e3)),m}function J(){return d(".timestamp").forEach(e=>{G(e)}),d(".reltime").forEach(e=>{S(e)}),[G,S]}function le(e,n){let t=0;return(...o)=>{t!==0&&window.clearTimeout(t),t=window.setTimeout(()=>{e(null,...o)},n)}}function ae(e,n,t,o,r){if(!e)return;let s=e.id+"-list",i=document.createElement("datalist"),a=document.createElement("option");a.value="",a.innerText="Loading...",i.appendChild(a),i.id=s,e.parentElement?.prepend(i),e.setAttribute("autocomplete","off"),e.setAttribute("list",s);let f={},m="";function c(l){let u=n;return u.includes("?")?u+"&t=json&"+t+"="+encodeURIComponent(l):u+"?t=json&"+t+"="+encodeURIComponent(l)}function g(l){let u=f[l];!u||!u.frag||(m=l,i.replaceChildren(u.frag.cloneNode(!0)))}function p(){let l=e.value;if(l.length===0)return;let u=c(l),w=!l||!m;if(!w){let h=f[m];h&&(w=!h.data.find(k=>l===r(k)))}if(w){if(f[l]&&f[l].url===u){g(l);return}fetch(u,{credentials:"include"}).then(h=>h.json()).then(h=>{if(!h)return;let k=Array.isArray(h)?h:[h],A=document.createDocumentFragment(),U=10;for(let M=0;M<k.length&&U>0;M++){let F=r(k[M]),ie=o(k[M]);if(F){let H=document.createElement("option");H.value=F,H.innerText=ie,A.appendChild(H),U--}}f[l]={url:u,data:k,frag:A,complete:!1},g(l)})}}e.oninput=le(p,250),console.log("managing ["+e.id+"] autocomplete")}function P(){return ae}function Y(){document.addEventListener("keydown",e=>{e.key==="Escape"&&document.location.hash.substring(0,7)==="#modal-"&&window.history.back()}),d(".backdrop, .modal-close").forEach(e=>{e.addEventListener("click",n=>{n.preventDefault(),window.history.back()})})}function N(e,n,t){return n||(n=18),t==null&&(t="icon"),`<svg class="${t}" style="width: ${n}px; height: ${n+"px"};"><use xlink:href="#svg-${e}"></use></svg>`}function de(e,n){return e.parentElement!==n.parentElement?null:e===n?0:e.compareDocumentPosition(n)&Node.DOCUMENT_POSITION_FOLLOWING?-1:1}var x;function D(e){let n=[],t=d(".item .value",e);for(let r of t)n.push(r.innerText);let o=T("input.result",e);o.value=n.join(", ")}function me(e){let n=e.parentElement?.parentElement;e.remove(),n&&D(n)}function Q(e){let n=T(".value",e),t=T(".editor",e);t.value=n.innerText;let o=()=>{if(t.value===""){e.remove();return}n.innerText=t.value,v(n,!0),v(t,!1);let r=e.parentElement?.parentElement;r&&D(r)};t.onblur=o,t.onkeydown=r=>{if(r.code==="Enter")return r.preventDefault(),o(),!1},v(n,!1),v(t,!0),t.focus()}function V(e,n){let t=document.createElement("div");t.className="item",t.draggable=!0,t.ondragstart=i=>{i.dataTransfer?.setDragImage(document.createElement("div"),0,0),t.classList.add("dragging"),x=t},t.ondragover=()=>{let i=de(t,x);if(!i)return;let a=i===-1?t:t.nextSibling;x.parentElement?.insertBefore(x,a),D(n)},t.ondrop=i=>{i.preventDefault()},t.ondragend=i=>{t.classList.remove("dragging"),i.preventDefault()};let o=document.createElement("div");o.innerText=e,o.className="value",o.onclick=()=>{Q(t)},t.appendChild(o);let r=document.createElement("input");r.className="editor",t.appendChild(r);let s=document.createElement("div");return s.innerHTML=N("times",13),s.className="close",s.onclick=()=>{me(t)},t.appendChild(s),t}function ue(e,n){let t=V("",n);e.appendChild(t),Q(t)}function K(e){let n=y("input.result",e);if(!n)return;let t=T(".tags",e),o=n.value.split(",").map(s=>s.trim()).filter(s=>s!=="");v(n,!1),t.innerHTML="";for(let s of o)t.appendChild(V(s,e));y(".add-item",e)?.remove();let r=document.createElement("div");r.className="add-item",r.innerHTML=N("plus",22),r.onclick=()=>{ue(t,e)},e.insertBefore(r,T(".clear",e))}function Z(){for(let e of d(".tag-editor"))K(e);return K}var X="--selected";function fe(e){let n=e.parentElement?.parentElement?.querySelector("input");if(!n)throw new Error("no associated input found");n.value="\u2205"}function O(e){e.onreset=()=>O(e);let n={},t={};for(let o of e.elements){let r=o;if(r.name.length>0)if(r.name.endsWith(X))t[r.name]=r;else{(r.type!=="radio"||r.checked)&&(n[r.name]=r.value);let s=()=>{let i=t[r.name+X];i&&(i.checked=n[r.name]!==r.value)};r.onchange=s,r.onkeyup=s}}}function z(){for(let e of d("form.editor"))O(e);return[fe,O]}var pe=[];function ee(e,n,t){let o=e.querySelectorAll(n);if(o.length===0)throw new Error("empty query selector ["+n+"]");o.forEach(r=>t(r))}function b(e,n,t){ee(e,n,o=>{o.style.backgroundColor=t})}function E(e,n,t){ee(e,n,o=>{o.style.color=t})}function ge(e,n,t){let o=document.querySelector("#mockup-"+e);if(!o){console.error("can't find mockup for mode ["+e+"]");return}switch(n){case"color-foreground":E(o,".mock-main",t);break;case"color-background":b(o,".mock-main",t);break;case"color-foreground-muted":E(o,".mock-main .mock-muted",t);break;case"color-background-muted":b(o,".mock-main .mock-muted",t);break;case"color-link-foreground":E(o,".mock-main .mock-link",t);break;case"color-link-visited-foreground":E(o,".mock-main .mock-link-visited",t);break;case"color-nav-foreground":E(o,".mock-nav",t),E(o,".mock-nav .mock-link",t);break;case"color-nav-background":b(o,".mock-nav",t);break;case"color-menu-foreground":E(o,".mock-menu",t),E(o,".mock-menu .mock-link",t);break;case"color-menu-background":b(o,".mock-menu",t);break;case"color-menu-selected-foreground":E(o,".mock-menu .mock-link-selected",t);break;case"color-menu-selected-background":b(o,".mock-menu .mock-link-selected",t);break;default:console.error("invalid key ["+n+"]")}}function te(){for(let e of d(".color-var")){let n=e.dataset.var,t=e.dataset.mode;pe.push(n),!(!n||n.length===0)&&(e.oninput=()=>{ge(t,n,e.value)})}}var ne=!1;function oe(e){if(e||(e="/connect"),e.indexOf("ws")===0)return e;let n=document.location,t="ws";return n.protocol==="https:"&&(t="wss"),e.indexOf("/")!==0&&(e="/"+e),t+`://${n.host}${e}`}var L=class{constructor(n,t,o,r,s){this.debug=n,this.open=t,this.recv=o,this.err=r,this.url=oe(s),this.connected=!1,this.pauseSeconds=1,this.pendingMessages=[],this.connect()}connect(){window.onbeforeunload=()=>{ne=!0},this.connectTime=Date.now(),this.sock=new WebSocket(oe(this.url));let n=this;this.sock.onopen=()=>{n.connected=!0,n.pendingMessages.forEach(n.send),n.pendingMessages=[],n.debug&&console.log("WebSocket connected"),n.open()},this.sock.onmessage=t=>{let o=JSON.parse(t.data);n.debug&&console.debug("[socket]: receive",o),o.cmd==="close-connection"?n.disconnect():n.recv(o)},this.sock.onerror=t=>()=>{n.err("socket",t.type)},this.sock.onclose=()=>{if(ne||n.closed)return;n.connected=!1;let t=n.connectTime?Date.now()-n.connectTime:0;t>0&&t<2e3?(n.pauseSeconds=n.pauseSeconds*2,n.debug&&console.debug(`socket closed immediately, reconnecting in ${n.pauseSeconds} seconds`),setTimeout(()=>{n.connect()},n.pauseSeconds*1e3)):(console.debug("socket closed after ["+t+"ms]"),setTimeout(()=>{n.connect()},n.pauseSeconds*500))}}disconnect(){this.closed=!0,setTimeout(()=>{this.sock?.close(),console.debug("[socket] closed")},500)}send(n){if(this.debug&&console.debug("out",n),!this.sock)throw new Error("socket not initialized");if(this.connected){let t=JSON.stringify(n,null,2);this.sock.send(t)}else this.pendingMessages.push(n)}};function re(){return L}function se(e,n,t,o,r){let s=()=>{e&&console.log("[socket]: open")},i=()=>{if(t){let g=document.createElement("tr"),p=document.createElement("th");p.innerText=n.children.length.toString();let l=document.createElement("td");return g.append(p,l),n.append(g),l}let c=document.createElement("div");return n.append(c),c},a=null,f=c=>{if(c.cmd!=="output"){if(r.length===0)console.log("unknown command ["+c.cmd+"] received");else for(let u of r)u(c);return}c.param.html===void 0&&console.log("no [html] key in message param: "+JSON.stringify(c,null,2));let g=c.param.html,p="";for(let u of g)a||(a=i()),u===`
`?(p===""&&(p="&nbsp;"),a.innerHTML+=p,p="",a=null):p+=u;a&&(a.innerHTML+=p);let l=T("#content");l.scrollTo(0,l.scrollHeight)},m=(c,g)=>{console.error("socket error",c,g)};return new L(!1,s,f,m,o)}function he(){let[e,n]=z(),[t,o]=J();window.loadtoad={wireTime:t,relativeTime:o,autocomplete:P(),setSiblingToNull:e,initForm:n,flash:_(),tags:Z(),Socket:re(),socketLog:se},q(),R(),j(),Y(),te(),window.audit=B}document.addEventListener("DOMContentLoaded",he);})();
//# sourceMappingURL=client.js.map
