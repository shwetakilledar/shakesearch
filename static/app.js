const Controller = {
	search: (ev) => {
		ev.preventDefault();
		const form = document.getElementById('form');
		const data = Object.fromEntries(new FormData(form));
		fetch(`/search?q=${data.query}`).then((response) => {
			response
				.json()
				.then((result) => {
					if (result.data && typeof result.data === 'string') {
						didYouMean.innerHTML =
							'Did you mean "' +
							result.data +
							'"?' +
							' Please update your search';
						Controller.updateTable([]);
					} else {
						Controller.updateTable(result.data);
					}
				})
				.catch((err) => {
					console.log('err', err);
				});
		});
	},

	updateTable: (results) => {
		const table = document.getElementById('table-body');
		const rows = [];
		for (let result of results) {
			rows.push(`<tr>${result}<tr/>`);
		}
		table.innerHTML = rows;
	},
};

const form = document.getElementById('form');
const didYouMean = document.getElementById('did-you-mean');
const correctedWord = document.getElementById('corrected-word');
form.addEventListener('submit', Controller.search);
