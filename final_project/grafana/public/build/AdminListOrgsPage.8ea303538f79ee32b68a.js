(window.webpackJsonp=window.webpackJsonp||[]).push([[9],{OQsP:function(e,t,s){"use strict";s.r(t);var n,c,i,r,a=s("q1tI"),l=s("lzJ5"),d=s("ZGyg"),j=s("/MKj"),o=s("kDLi"),b=s("t8hP"),h=s("nKUr");const O=({orgs:e,onDelete:t})=>{const[s,l]=Object(a.useState)();return Object(h.jsxs)("table",{className:"filter-table form-inline filter-table--hover",children:[Object(h.jsx)("thead",{children:Object(h.jsxs)("tr",{children:[n||(n=Object(h.jsx)("th",{children:"ID"})),c||(c=Object(h.jsx)("th",{children:"Name"})),Object(h.jsx)("th",{style:{width:"1%"}})]})}),Object(h.jsx)("tbody",{children:e.map(e=>Object(h.jsxs)("tr",{children:[Object(h.jsx)("td",{className:"link-td",children:Object(h.jsx)("a",{href:"admin/orgs/edit/"+e.id,children:e.id})}),Object(h.jsx)("td",{className:"link-td",children:Object(h.jsx)("a",{href:"admin/orgs/edit/"+e.id,children:e.name})}),Object(h.jsx)("td",{className:"text-right",children:Object(h.jsx)(o.Button,{variant:"destructive",size:"sm",icon:"times",onClick:()=>l(e)})})]},`${e.id}-${e.name}`))}),s&&Object(h.jsx)(o.ConfirmModal,{isOpen:!0,icon:"trash-alt",title:"Delete",body:Object(h.jsxs)("div",{children:["Are you sure you want to delete '",s.name,"'?",i||(i=Object(h.jsx)("br",{}))," ",r||(r=Object(h.jsx)("small",{children:"All dashboards for this organization will be removed!"}))]}),confirmText:"Delete",onDismiss:()=>l(void 0),onConfirm:()=>{t(s.id),l(void 0)}})]})};var x,g=s("51gB"),m=s.n(g);s.d(t,"AdminListOrgsPages",(function(){return u}));const u=()=>{const e=Object(j.useSelector)(e=>e.navIndex),t=Object(l.a)(e,"global-orgs"),[s,n]=m()(async()=>await(async()=>await Object(b.getBackendSrv)().get("/api/orgs"))(),[]);return Object(a.useEffect)(()=>{n()},[n]),Object(h.jsx)(d.b,{navModel:t,children:Object(h.jsx)(d.b.Contents,{children:Object(h.jsxs)(h.Fragment,{children:[x||(x=Object(h.jsxs)("div",{className:"page-action-bar",children:[Object(h.jsx)("div",{className:"page-action-bar__spacer"}),Object(h.jsx)(o.LinkButton,{icon:"plus",href:"org/new",children:"New org"})]})),s.loading&&"Fetching organizations",s.error,s.value&&Object(h.jsx)(O,{orgs:s.value,onDelete:e=>{(async e=>await Object(b.getBackendSrv)().delete("/api/orgs/"+e))(e).then(()=>n())}})]})})})};t.default=u}}]);
//# sourceMappingURL=AdminListOrgsPage.8ea303538f79ee32b68a.js.map