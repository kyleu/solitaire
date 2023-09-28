"use strict";(()=>{function d(e,n){let t;n?t=n.querySelectorAll(e):t=document.querySelectorAll(e);let r=[];return t.forEach(o=>{r.push(o)}),r}function k(e,n){let t=d(e,n);switch(t.length){case 0:return;case 1:return t[0];default:console.warn(`found [${t.length}] elements with selector [${e}], wanted zero or one`)}}function p(e,n){let t=k(e,n);if(!t)throw new Error(`no element found for selector [${e}]`);return t}function F(e,n){return typeof e=="string"&&(e=p(e)),e.innerHTML=n,e}function E(e,n,t="block"){return typeof e=="string"&&(e=p(e)),e.style.display=n?t:"none",e}function J(e,n,...t){let r=document.createElement(e);for(let o in n)if(o==="for"&&(o="for"),o==="className"&&(o="class"),o&&n.hasOwnProperty(o)){let i=n[o];o==="dangerouslySetInnerHTML"?F(r,i.__html):i===!0?r.setAttribute(o,o):i!==!1&&i!==null&&i!==void 0&&r.setAttribute(o,i.toString())}for(let o of t)if(Array.isArray(o))o.forEach(i=>{if(o==null)throw new Error(`child array for tag [${e}] is ${o}
${r.outerHTML}`);if(i==null)throw new Error(`child for tag [${e}] is ${i}
${r.outerHTML}`);typeof i=="string"&&(i=document.createTextNode(i)),r.appendChild(i)});else{if(o==null)throw new Error(`child for tag [${e}] is ${o}
${r.outerHTML}`);o.nodeType||(o=document.createTextNode(o.toString())),r.appendChild(o)}return r}function ue(e,...n){let t=document.createElement("li");t.innerText=e;for(let r of n){let o=document.createElement("pre");typeof r=="string"?o.innerText=r:o.innerText=JSON.stringify(r,null,2),t.appendChild(o)}return t}function q(e,...n){let t=k("#audit-log");t?t.appendChild(ue(e,...n)):console.log(`### Audit ###
`+e,...n)}function R(){for(let e of d(".menu-container .final"))e.scrollIntoView({block:"center"})}var x="mode-light",w="mode-dark";function W(){for(let e of d(".mode-input"))e.onclick=()=>{switch(e.value){case"":document.body.classList.remove(x),document.body.classList.remove(w);break;case"light":document.body.classList.add(x),document.body.classList.remove(w);break;case"dark":document.body.classList.remove(x),document.body.classList.add(w);break}}}function j(e){setTimeout(()=>{e.style.opacity="0",setTimeout(()=>e.remove(),500)},5e3)}function _(e,n,t){let r=document.getElementById("flash-container");r===null&&(r=document.createElement("div"),r.id="flash-container",document.body.insertAdjacentElement("afterbegin",r));let o=document.createElement("div");o.className="flash";let i=document.createElement("input");i.type="radio",i.style.display="none",i.id="hide-flash-"+e,o.appendChild(i);let s=document.createElement("label");s.htmlFor="hide-flash-"+e;let c=document.createElement("span");c.innerHTML="\xD7",s.appendChild(c),o.appendChild(s);let l=document.createElement("div");l.className="content flash-"+n,l.innerText=t,o.appendChild(l),r.appendChild(o),j(o)}function B(){let e=document.getElementById("flash-container");if(e===null)return _;let n=e.querySelectorAll(".flash");if(n.length>0)for(let t of n)j(t);return _}function G(){for(let e of d(".link-confirm"))e.onclick=()=>{let n=e.dataset.message;return n&&n.length===0&&(n="Are you sure?"),confirm(n)}}function H(e,n){let t=(e||"").replace(/-/gu,"/").replace(/[TZ]/gu," ")+" UTC",r=new Date(t),o=(new Date().getTime()-r.getTime())/1e3,i=Math.floor(o/86400),s=r.getFullYear(),c=r.getMonth()+1,l=r.getDate();if(isNaN(i)||i<0||i>=31)return s.toString()+"-"+(c<10?"0"+c.toString():c.toString())+"-"+(l<10?"0"+l.toString():l.toString());let a,m;return i===0?o<5?(m=1,a="just now"):o<60?(m=1,a=Math.floor(o)+" seconds ago"):o<120?(m=10,a="1 minute ago"):o<3600?(m=30,a=Math.floor(o/60)+" minutes ago"):o<7200?(m=60,a="1 hour ago"):(m=60,a=Math.floor(o/3600)+" hours ago"):i===1?(m=600,a="yesterday"):i<7?(m=600,a=i+" days ago"):(m=6e3,a=Math.ceil(i/7)+" weeks ago"),n&&(n.innerText=a,setTimeout(()=>H(e,n),m*1e3)),a}function X(){return d(".reltime").forEach(e=>{H(e.dataset.time||"",e)}),H}function me(e,n){let t=0;return(...r)=>{t!==0&&window.clearTimeout(t),t=window.setTimeout(()=>{e(null,...r)},n)}}function fe(e,n,t,r,o){var A;if(!e)return;let i=e.id+"-list",s=document.createElement("datalist"),c=document.createElement("option");c.value="",c.innerText="Loading...",s.appendChild(c),s.id=i,(A=e.parentElement)==null||A.prepend(s),e.setAttribute("autocomplete","off"),e.setAttribute("list",i);let l={},a="";function m(u){let f=n;return f.includes("?")?f+"&t=json&"+t+"="+encodeURIComponent(u):f+"?t=json&"+t+"="+encodeURIComponent(u)}function D(u){let f=l[u];!f||!f.frag||(a=u,s.replaceChildren(f.frag.cloneNode(!0)))}function ae(){let u=e.value;if(u.length===0)return;let f=m(u),L=!u||!a;if(!L){let g=l[a];g&&(L=!g.data.find(T=>u===o(T)))}if(L){if(l[u]&&l[u].url===f){D(u);return}fetch(f,{credentials:"include"}).then(g=>g.json()).then(g=>{if(!g)return;let T=Array.isArray(g)?g:[g],O=document.createDocumentFragment(),$=10;for(let y=0;y<T.length&&$>0;y++){let U=o(T[y]),de=r(T[y]);if(U){let b=document.createElement("option");b.value=U,b.innerText=de,O.appendChild(b),$--}}l[u]={url:f,data:T,frag:O,complete:!1},D(u)})}}e.oninput=me(ae,250),console.log("managing ["+e.id+"] autocomplete")}function P(){return fe}function V(){document.addEventListener("keydown",e=>{e.key==="Escape"&&document.location.hash.startsWith("#modal-")&&(document.location.hash="")})}function I(e,n,t){return n||(n=18),t==null&&(t="icon"),`<svg class="${t}" style="width: ${n}px; height: ${n+"px"};"><use xlink:href="#svg-${e}"></use></svg>`}function pe(e,n){return e.parentElement!==n.parentElement?null:e===n?0:e.compareDocumentPosition(n)&Node.DOCUMENT_POSITION_FOLLOWING?-1:1}var M;function C(e){let n=[],t=d(".item .value",e);for(let o of t)n.push(o.innerText);let r=p("input.result",e);r.value=n.join(", ")}function ge(e){var t;let n=(t=e.parentElement)==null?void 0:t.parentElement;e.remove(),n&&C(n)}function K(e){let n=p(".value",e),t=p(".editor",e);t.value=n.innerText;let r=()=>{var i;if(t.value===""){e.remove();return}n.innerText=t.value,E(n,!0),E(t,!1);let o=(i=e.parentElement)==null?void 0:i.parentElement;o&&C(o)};t.onblur=r,t.onkeydown=o=>{if(o.code==="Enter")return o.preventDefault(),r(),!1},E(n,!1),E(t,!0),t.focus()}function Q(e,n){let t=document.createElement("div");t.className="item",t.draggable=!0,t.ondragstart=s=>{var c;(c=s.dataTransfer)==null||c.setDragImage(document.createElement("div"),0,0),t.classList.add("dragging"),M=t},t.ondragover=()=>{var l;let s=pe(t,M);if(!s)return;let c=s===-1?t:t.nextSibling;(l=M.parentElement)==null||l.insertBefore(M,c),C(n)},t.ondrop=s=>{s.preventDefault()},t.ondragend=s=>{t.classList.remove("dragging"),s.preventDefault()};let r=document.createElement("div");r.innerText=e,r.className="value",r.onclick=()=>{K(t)},t.appendChild(r);let o=document.createElement("input");o.className="editor",t.appendChild(o);let i=document.createElement("div");return i.innerHTML=I("times",13),i.className="close",i.onclick=()=>{ge(t)},t.appendChild(i),t}function he(e,n){let t=Q("",n);e.appendChild(t),K(t)}function Y(e){var i;let n=p("input.result",e),t=p(".tags",e),r=n.value.split(",").map(s=>s.trim()).filter(s=>s!=="");E(n,!1),t.innerHTML="";for(let s of r)t.appendChild(Q(s,e));(i=k(".add-item",e))==null||i.remove();let o=document.createElement("div");o.className="add-item",o.innerHTML=I("plus",22),o.onclick=()=>{he(t,e)},e.insertBefore(o,p(".clear",e))}function Z(){for(let e of d(".tag-editor"))Y(e);return Y}function Te(e){let n=document.createElement("thead"),t=document.createElement("tr"),r=!0;return e.forEach(o=>{let i=document.createElement("th");r&&(i.classList.add("shrink"),r=!1),i.innerText=o.title,t.appendChild(i)}),n.appendChild(t),n}function Ee(e,n){let t=document.createElement("tr"),r=[];return e.forEach(o=>{let i=document.createElement("td");r.push(i);let s=n[o.key];if(s==null){let c=document.createElement("em");c.innerText="-",i.appendChild(c)}else if(o.type==="code"){let c=document.createElement("pre");c.innerText=JSON.stringify(s,null,2),i.appendChild(c)}else i.innerText=s.toString();t.appendChild(i)}),[t,r]}function z(e,n){let t=document.createElement("table"),r=[];t.classList.add("min-200"),t.appendChild(Te(e));let o=document.createElement("tbody");n.forEach(s=>{let[c,l]=Ee(e,s);o.appendChild(c),r.push(l)}),t.appendChild(o);let i=document.createElement("div");return i.classList.add("overflow"),i.classList.add("full-width"),i.appendChild(t),i}function ke(e){var c,l;let n=(c=e.dataset.key)!=null?c:"editor",t=(l=e.dataset.columns)!=null?l:"[]",r=JSON.parse(t.replace(/\\"/gu,'"')),o=p("textarea",e),i=JSON.parse(o.value);if(i==null&&(i=[]),!Array.isArray(i))throw new Error("input value for element ["+n+"] of type ["+typeof i+"] must be an array");let s=z(r,i);d(".toggle-editor-"+n).forEach(a=>{a.innerText="Edit",a.onclick=()=>{a.innerText==="Edit"?(a.innerText="View",s.remove(),o.hidden=!1):(a.innerText="Edit",s.remove(),i=JSON.parse(o.value),i==null&&(i=[]),s=z(r,i),e.appendChild(s),o.hidden=!0)}}),e.appendChild(s),o.hidden=!0}function ee(){d(".rich-editor").forEach(e=>{ke(e)})}var te="--selected";function ve(e){var t,r;let n=(r=(t=e.parentElement)==null?void 0:t.parentElement)==null?void 0:r.querySelector("input");if(!n)throw new Error("no associated input found");n.value="\u2205"}function S(e){e.onreset=()=>S(e);let n={},t={};for(let r of e.elements){let o=r;if(o.name.length>0)if(o.name.endsWith(te))t[o.name]=o;else{(o.type!=="radio"||o.checked)&&(n[o.name]=o.value);let i=()=>{let s=t[o.name+te];s&&(s.checked=n[o.name]!==o.value)};o.onchange=i,o.onkeyup=i}}}function ne(){for(let e of d("form.editor"))S(e);return[ve,S]}var ye=[];function oe(e,n,t){let r=e.querySelectorAll(n);if(r.length===0)throw new Error("empty query selector ["+n+"]");r.forEach(o=>t(o))}function v(e,n,t){oe(e,n,r=>{r.style.backgroundColor=t})}function h(e,n,t){oe(e,n,r=>{r.style.color=t})}function Me(e,n,t){let r=document.querySelector("#mockup-"+e);if(!r){console.error("can't find mockup for mode ["+e+"]");return}switch(n){case"color-foreground":h(r,".mock-main",t);break;case"color-background":v(r,".mock-main",t);break;case"color-foreground-muted":h(r,".mock-main .mock-muted",t);break;case"color-background-muted":v(r,".mock-main .mock-muted",t);break;case"color-link-foreground":h(r,".mock-main .mock-link",t);break;case"color-link-visited-foreground":h(r,".mock-main .mock-link-visited",t);break;case"color-nav-foreground":h(r,".mock-nav",t),h(r,".mock-nav .mock-link",t);break;case"color-nav-background":v(r,".mock-nav",t);break;case"color-menu-foreground":h(r,".mock-menu",t),h(r,".mock-menu .mock-link",t);break;case"color-menu-background":v(r,".mock-menu",t);break;case"color-menu-selected-foreground":h(r,".mock-menu .mock-link-selected",t);break;case"color-menu-selected-background":v(r,".mock-menu .mock-link-selected",t);break;default:console.error("invalid key ["+n+"]")}}function re(){for(let e of d(".color-var")){let n=e.dataset.var,t=e.dataset.mode;ye.push(n),!(!n||n.length===0)&&(e.oninput=()=>{Me(t,n,e.value)})}}var ie=!1;function se(e){if(e||(e="/connect"),e.indexOf("ws")===0)return e;let n=document.location,t="ws";return n.protocol==="https:"&&(t="wss"),e.indexOf("/")!==0&&(e="/"+e),t+`://${n.host}${e}`}var N=class{constructor(n,t,r,o,i){this.debug=n,this.open=t,this.recv=r,this.err=o,this.url=se(i),this.connected=!1,this.pauseSeconds=1,this.pendingMessages=[],this.connect()}connect(){window.onbeforeunload=()=>{ie=!0},this.connectTime=Date.now(),this.sock=new WebSocket(se(this.url));let n=this;this.sock.onopen=()=>{n.connected=!0,n.pendingMessages.forEach(n.send),n.pendingMessages=[],n.debug&&console.log("WebSocket connected"),n.open()},this.sock.onmessage=t=>{let r=JSON.parse(t.data);n.debug&&console.debug("[socket]: receive",r),n.recv(r)},this.sock.onerror=t=>()=>{n.err("socket",t.type)},this.sock.onclose=()=>{if(ie)return;n.connected=!1;let t=n.connectTime?Date.now()-n.connectTime:0;t>0&&t<2e3?(n.pauseSeconds=n.pauseSeconds*2,n.debug&&console.debug(`socket closed immediately, reconnecting in ${n.pauseSeconds} seconds`),setTimeout(()=>{n.connect()},n.pauseSeconds*1e3)):(console.debug("socket closed after ["+t+"ms]"),setTimeout(()=>{n.connect()},n.pauseSeconds*500))}}disconnect(){}send(n){if(this.debug&&console.debug("out",n),!this.sock)throw new Error("socket not initialized");if(this.connected){let t=JSON.stringify(n,null,2);this.sock.send(t)}else this.pendingMessages.push(n)}};function ce(){return N}function le(){console.log("!!!")}function Le(){le()}window.gameInit=Le;function be(){let[e,n]=ne();window.solitaire={relativeTime:X(),autocomplete:P(),setSiblingToNull:e,initForm:n,flash:B(),tags:Z(),Socket:ce()},R(),W(),G(),V(),re(),ee(),window.audit=q,window.JSX=J}document.addEventListener("DOMContentLoaded",be);})();
//# sourceMappingURL=client.js.map
