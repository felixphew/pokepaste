const makeListener = (tgt, id, href) => {
	tgt.addEventListener("change", e => {
		const checked = e.currentTarget.checked;
		const cssid = id + "css";
		let css = document.getElementById(cssid);
		if (checked) {
			if (!css) {
				css = document.createElement("link");
				css.id = cssid;
				css.rel = "stylesheet";
				css.href = href;
				document.head.appendChild(css);
			}
		} else {
			if (css) {
				css.parentNode.removeChild(css);
			}
		}

		let date = new Date();
		date.setTime(date.getTime() + 31536000);

		document.cookie = `${id}=${checked}; max-age=31536000`;
	});

	tgt.checked = false;

	if (document.cookie.includes(`${id}=true`)) tgt.click();
};
	
makeListener(document.getElementById("vgcmode"), "vgc", "/css/paste-vgc.css");
makeListener(document.getElementById("evivmode"), "eviv", "/css/syntax-eviv.css");
makeListener(document.getElementById("lightmode"), "light", "/css/paste-light.css");
