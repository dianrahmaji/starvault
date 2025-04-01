const path = require('path');

module.exports = {
	'!./(apps|packages)/**/*': ['prettier --write'],
	'./(apps|packages)/*/src/**/(*.{js,ts,json,yaml})': async (stagedFiles) => {
		const projectNames = [
			...new Set(stagedFiles.map((file) => file.split(path.sep)[1])),
		];

		return [
			`prettier --write ${stagedFiles.join(' ')}`,
			...projectNames.map(() => `turbo run lint -- --fix`),
		];
	},
};
