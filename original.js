
function call(){
	const randomStr = function () {
		return window.crypto.randomUUID().toString().substring(0, 8).toUpperCase();
	};

	/** some comment */
	alert(randomStr());
}