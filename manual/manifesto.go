package manual

func init() {
	Register("manifesto", manifesto)
}

func manifesto() string {
	return `
Scarlet was built on the following ideas and principles:

	1. Soft-Magic Themed

		Scarlet is a soft-magic themed tool. Programming is a kind of hard-magic
		with an unrelenting vortex of soul sucking rationalism veiled beneath
		its technoshiny exterior --probably why women aren't that interested--.
		While it's all for the best, I found the initial enchantment wears off
		after a few years in the slave pits. Odd thing really, because
		programming is filled with things that don't make sense unless you were
		in the right place, at the right time, and inside the right mind; even
		then, only if you're lucky.

		I wanted to inject some true magic. Magic that is shaped by the minds
		of practictioners without compromising derived solutions. Magic that
		cannot be rationalised yet does not need to be. Magic that wears a
		pointy hat, carry's a dai-katana, and talks with a feminine irish
		accent.

	2. No dependencies

		Scrolls (Scarlet scripts) have no native way to import other scrolls.
		Scarlet was intended as a secretary language to complete simple, yet
		essential, quests such as building applications, controlling pipelines,
		and a customisable replacement for bash scripts. Avoiding dependencies
		has the nice effect of avoiding a great source of complexity and the
		inappropriate use as a systems language.

	3. Easy integration

		Scarlet emphasises the creation of spells (inbuilt functions) for
		generic functionality. Spells are written in the underlying systems
		language, e.g. Go, so their external libraries may be used. Simply
		register the spell and recompile Scarlet. I envisioned a user, team,
		or commmunity will take a copy of the base tool then populate it with
		domain or team specific spells using patterns befitting use cases or
		the authors biases.

	4. Context specific

		Contrary to modern scripting tools, Scarlet scrolls are designed to be
		platform and situation specific, that is, scripts are written for a
		single purpose and usually for a single target platform. This may seem
		rather restrictive but it's to encourage context driven solutions and
		surpress the compelling urge to abstract everything. If you can't live
		with that you can always create a spell to import functions or something.

		"THAT'S MORTALS FOR YOU. THEY'VE ONLY GOT A FEW YEARS IN THIS WORLD AND
		THEY SPEND THEM ALL IN MAKING THINGS COMPLICATED FOR THEMSELVES."
			- 'Mort' by Terry Pratchett

	5. Minimalist

		Scarlet favours spells over native syntax, vis only the bare necessities
		form the base language, everything else is better off as a spell that
		can be easily modified. This also makes the rare addition of new native
		features a breeze.

		"Take it from me, there's nothing more terrible than someone out to do
		the world a favour."
			- 'Sourcery' by Terry Pratchett

	6. Light and portable

		The Scarlet executable will hopefully be light, portable, and require
		no installation process; much like Lua. With time and hope a Rust
		implementation will be built precisely for embedding in other programs
		and repositories.

		"'What's a philosopher?' said Brutha. 'Someone who's bright enough to
		find a job with no heavy lifting,' said a voice in his head."
			- 'Small Gods' by Terry Pratchett

Use cases:

	I intended for a very small binary so I could include it within	code
	repositories. Rust would have been a better choice for this
	optimisation but I decided to build an easier Go version first to	try
	out the idea and learn how to parse code. It could then be used to
	build and run applications both within pipelines and workstations
	without requiring additional tools be installed into container images
	and the such like. I want to automate without installation pains.

	With this I can create language independent Web API testing scripts
	so I can more easily switch a web server's implementation language
	and avoid self inflicted vendor lock in. Current tools were either
	too heavy weight or painfully complex. Project building,
	configuration, and deployment are another activity I want more
	control over.

	I also wanted to do general purpose scripting. There are plenty of
	languages that can assist with this but I really craved specific tools
	free of dependencies. I wanted to be able to change the langauge
	each time I noticed it was woefully incapable of satisfying me.

No, just no! I'm strongly for fitting the tool to the job and not the
other way around so here are a few use cases I do not recommend using
Scarlet for:

	Backend web programming. That's best left to systems tools such as Go,
	Java, C#, etc. However, I do intend to create spells for quickly serving
	static content and file storage on a local network.

	Anything that needs to scale or use concurrency. Again Go, Rust, and
	many JVM languages are good choices.

	Maths, science, or running numeric algorithms. That's best left to tools
	like R or library rich glue languages like Python.

	Avoid using it for critical systems! I wrote the code for me and don't
	want innocent bystanders (if such people exist) getting hurt.
	
	"A catastrophe curve, Mr. Bucket, is what Software runs along. Software
	happens because a large number of things amazingly fail without quite
	sinking projects, Mr. Bucket. It works because of hatred and love and
	nerves. All the time. This isn't cheese. This is Software. If you
	wanted a quiet retirement, Mr. Bucket, you shouldn't have bought the
	Software House.	You should have done something peaceful, like alligator
	dentistry."
		- (Original) 'Maskerade' by Terry Pratchett
		- Adapted to context by Paulio`
}
