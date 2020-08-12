package manual

func init() {
	Register("", overview)
	Register("overview", overview)
	Register("comment", comments)
	Register("comments", comments)
}

func overview() string {
	return `
Scarlet's language documentation.

Usage:

	scarlet docs [search term]

Search terms:

	spell              What are spells and how to use them?
	spells             List spells
	variables          How to use variables?
	comments           Writing comments
	types              Variable types and their uses
	<spell-name>       Documentation for a specific spell, e.g. '@Println'

Scarlet:

	"Sometimes it's better to light a flamethrower than curse the darkness."
		- 'Men at Arms' by Terry Pratchett

	Scarlet is a template for building domain or environment specific
	scripting	tools. Think of it as a bash replacement and a Lua alternative
	but the source code is easy to modify and compile almost anywhere. First
	I will present the desirable characteristics that guide development, then
	a number of envisioned use cases. This should provide a good feel for why
	Scarlet	is being built and its intended purposes.

	1. Discworld Themed

		Scarlet is a discworld themed tool, because the logical domain that is
		programming is filled with things that don't make sense unless you were
		in the right place, at the right time, and inside the right head.
		There's also a servere lack of true magic, probably why the ladies
		arn't that interested.

	2. No dependencies

		Scarlett scripts have no native way to import other Scarlett scripts.
		This avoids the	many considerations and issues associated with the
		practice. Scarlet is designed specifcally for small problems and
		networkless environments.

		"Darkness isn't the opposite of light, it is simply its absence."
			- 'The Light Fantastic' by Terry Pratchett

	3. Easy integration

		Scarlet emphasises the creation of spells (inbuilt functions) for
		generic functionality. Spells are written in Go so external Go
		libraries can be used. Simply register the spell and recompile Scarlet.
		I envisioned the tool will be copied and then populated with domain or
		environment specific spells using any patterns the authors see fit.

		"The three rules of the Librarians of Time and Space are:
		1) Silence;
		2) Books must be returned no later than the date last shown; and
		3) Do not interfere with the nature of causality."
			- 'Guards! Guards!' by Terry Pratchett

	4. Context specific

		Unlike some modern scripting languages, Scarlett scripts are designed
		to be platform and situation specific, that is, scripts are written for
		a single purpose and against a specific version of the tool. This may
		seem rather restrictive but its to encourage context driven solutions
		and surpress the compelling urge to abstract everything.

		"THAT'S MORTALS FOR YOU. THEY'VE ONLY GOT A FEW YEARS IN THIS WORLD AND
		THEY SPEND THEM ALL IN MAKING THINGS COMPLICATED FOR THEMSELVES."
			- 'Mort' by Terry Pratchett

	5. Minimalist

		Scarlet favours spells over native syntax vis, if a feature is not used
		constantly or is niche then its better off as a spell that can more
		easily be modified. Fewer default native features allows uselful ones
		to be added quickly when desired.

		"Take it from me, there's nothing more terrible than someone out to do
		the world a favour."
			- 'Sourcery' by Terry Pratchett

	6. Customisable

		If you don't like the names of spells, change them.
		If you don't like the language keywords, change them.
		If you don't like function brackets, change them.

		"What don’t die can’t live.
		What don’t live can’t change.
		What don’t change can’t learn."
			- 'Lords and Ladies' by Terry Pratchett

	7. Light and portable

		The Scarlet executable is light, portable, and does not require an
		installation process; much like Lua. With time and hope a Rust
		implementation will be built precisely for embedding in other programs
		and repositories.

		"'What's a philosopher?' said Brutha. 'Someone who's bright enough to
		find a job with no heavy lifting,' said a voice in his head."
			- 'Small Gods' by Terry Pratchett

Good use cases:

	I intended for a very small binary so I could include it within	code
	repositories. Rust would have been a better choice for this
	optimisation but I decided to build an easier Go version first to	try
	out the idea and learn how to parse code. Once embedded within a
	repository it could be used to build and run applications both within
	pipelines and workstations without additional tools; the tools usually
	involve some god awful installation process.

	With this I could create language independent Web API testing scripts
	so I can more easily switch a web server's implementation language
	and avoid self inflicted vendor lock in. Current tools were either
	too heavy weight or painfully complex. Project building,
	configuration, and deployment was another activity I wanted more
	control over.

	I also wanted do general purpose scripting. There are plenty of
	languages that can assist with this but I really craved specific tools
	free of dependencies. I wanted to be able to change the langauge
	each time I noticed it was woefully incapable of satisfying me.

Bad use cases:

	I'm strongly for fitting the tool to the job and not the other way
	around so here are a few use cases that I recommend Scarlet not be
	used for.

	Scarlet is not intended, nor designed, for backend web programming.
	That's best left to more rigorous and much better supported tools such
	as Go, Java, and C#. However, I do intend to create spells for quickly
	serving static content and file storage on local networks.

	Anything that needs to scale or use concurrency. Again Go, Rust, and
	many JVM languages are good choices.

	It is not intended for maths, science, or running numeric algorithms.
	That's best left to tools like R or library rich glue languages like
	Python.

	Avoid using it for critical systems! I wrote the code for me and don't
	want innocent bystanders (if such people exist) getting hurt.
	
	"A catastrophe curve, Mr. Bucket, is what Software runs along. Software
	happens because a large number of things amazingly fail without quite
	sinking their project,	Mr. Bucket. It works because of hatred and love
	and nerves. All the	time. This isn't cheese. This is Software. If you
	wanted a quiet retirement, Mr. Bucket, you shouldn't have bought the
	Software House.	You should have done something peaceful, like alligator
	dentistry."
		- (Original) 'Maskerade' by Terry Pratchett
		- Adapted by Paulio`
}

func comments() string {
	return `
Comments provide a way for programmers to communicate to readers what they
think their code does in a manner that is easily misinterpretted. It is
customary to write comments whenever you feel like it and not based on the
ambiguaity or complexity of the functionality under comment.

Examples:

	# Comments start with the a pound symbol '#', sometimes referred to as hash,
	# and terminate at the end of the line. What you write in the comment is
	# entirely up to you.

	# This function adds two numbers
	add := F(a, b -> c) {
		c := a + b # The answer is stored in the return variable 'c'
	}`
}
