(window.webpackJsonp=window.webpackJsonp||[]).push([[44],{"0jq5":function(e,t,n){"use strict";n.d(t,"a",(function(){return o}));var a=n("q1tI"),i=n("EKT6"),r=n("kDLi"),s=n("nKUr");function l(){return(l=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var a in n)Object.prototype.hasOwnProperty.call(n,a)&&(e[a]=n[a])}return e}).apply(this,arguments)}class o extends a.PureComponent{render(){const{searchQuery:e,linkButton:t,setSearchQuery:n,target:a,placeholder:o="Search by name or type"}=this.props,c={href:null==t?void 0:t.href};return a&&(c.target=a),Object(s.jsxs)("div",{className:"page-action-bar",children:[Object(s.jsx)("div",{className:"gf-form gf-form--grow",children:Object(s.jsx)(i.a,{value:e,onChange:n,placeholder:o})}),t&&Object(s.jsx)(r.LinkButton,l({},c,{children:t.title}))]})}}},KuTy:function(e,t,n){"use strict";n.r(t),n.d(t,"PlaylistPage",(function(){return y})),n.d(t,"StartModal",(function(){return O}));var a,i=n("q1tI"),r=n("/MKj"),s=n("Obii"),l=n("ZGyg"),o=n("lzJ5"),c=n("Y8YH"),d=n("t8hP"),u=n("kDLi"),j=n("HJRA"),p=n("0jq5"),b=n("QQVG"),h=n("nKUr");const y=({navModel:e})=>{const[t,n]=Object(i.useState)(""),[r,s]=Object(i.useState)(),{value:o,loading:y}=Object(c.a)(async()=>Object(d.getBackendSrv)().get("/api/playlists",{query:t})),g=o&&o.length>0;let v=a||(a=Object(h.jsx)(b.a,{title:"There are no playlists created yet",buttonIcon:"plus",buttonLink:"playlists/new",buttonTitle:"Create Playlist",proTip:"You can use playlists to cycle dashboards on TVs without user control",proTipLink:"http://docs.grafana.org/reference/playlist/",proTipLinkTitle:"Learn more",proTipTarget:"_blank"}));return g&&(v=Object(h.jsx)(h.Fragment,{children:o.map(e=>Object(h.jsx)(u.Card,{heading:e.name,children:Object(h.jsxs)(u.Card.Actions,{children:[Object(h.jsx)(u.Button,{variant:"secondary",icon:"play",onClick:()=>s(e),children:"Start playlist"}),j.b.isEditor&&Object(h.jsx)(u.LinkButton,{variant:"secondary",href:"/playlists/edit/"+e.id,icon:"cog",children:"Edit playlist"},"edit")]})},e.id.toString()))})),Object(h.jsx)(l.b,{navModel:e,children:Object(h.jsxs)(l.b.Contents,{isLoading:y,children:[g&&Object(h.jsx)(p.a,{searchQuery:t,linkButton:{title:"New playlist",href:"/playlists/new"},setSearchQuery:n}),v,r&&Object(h.jsx)(O,{playlist:r,onDismiss:()=>s(void 0)})]})})};t.default=Object(r.connect)(e=>({navModel:Object(o.a)(e.navIndex,"playlists")}))(y);const O=({playlist:e,onDismiss:t})=>{const[n,a]=Object(i.useState)(!1),[r,l]=Object(i.useState)(!1);return Object(h.jsxs)(u.Modal,{isOpen:!0,icon:"play",title:"Start playlist",onDismiss:t,children:[Object(h.jsxs)(u.VerticalGroup,{children:[Object(h.jsx)(u.Field,{label:"Mode",children:Object(h.jsx)(u.RadioButtonGroup,{value:n,options:[{label:"Normal",value:!1},{label:"TV",value:"tv"},{label:"Kiosk",value:!0}],onChange:a})}),Object(h.jsx)(u.Checkbox,{label:"Autofit",description:"Panel heights will be adjusted to fit screen size",name:"autofix",value:r,onChange:e=>l(e.currentTarget.checked)})]}),Object(h.jsx)(u.Modal.ButtonRow,{children:Object(h.jsxs)(u.Button,{variant:"primary",onClick:()=>{const t={};n&&(t.kiosk=n),r&&(t.autofitpanels=!0),d.locationService.push(s.urlUtil.renderUrl("/playlists/play/"+e.id,t))},children:["Start ",e.name]})})]})}}}]);
//# sourceMappingURL=PlaylistPage.8ea303538f79ee32b68a.js.map