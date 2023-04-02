# Introduction

Yo, what up! if you're reading this then I'm super psyched because that means that you're thinking about contributing to Buster! Thanks so much for your time and consideration. It's rad people like you that make Buster such a useful CI/CD tool

I wrote this contributor's guide to help newcomers feel welcome. Getting started with a new project can be complicated and I wanted to make it as easy as possible for you to contribute and as easy as possible for me to help.

Currently any sincere pull request is a good request.
Buster is still in pre-release so there are so many way to contribute!
Here's a list of ideas but feel free to suggest anything I may have forgotten to include.

* Feature requests
* Unit and integration tests.
* Writing, editing, and translating tutorials, documentation, or blog posts.
* Auditing for accessibility.
* Bug reports.
* Bug triaging.
* Community management.
* Art! Dreams! Your excellence!
* Code that can be pulled into Buster itself.

# Contributor guidelines
### Excellence, and the contributor's code of conduct

First up, most importantly we have a contributor's code of conduct. For some reason the internet is a dehumanizing experience and it's easy to forget that aside from the bots we're all humans on this thing. Approach each other with kindness. Please read our [contributor's code of conduct](CODE_OF_CONDUCT.md) and when in doubt just remember our one true rule as once spoken by the ever so wise duo of Bill and Ted.

`Bill: Be excellent, ... to each other, ...`

`Ted: and party on, dudes! [sic]`

![Abraham Licoln saying, "Be Excellent to each other and party on dudes!". [sic]](https://media.giphy.com/media/ef0zYcF7AKu4b0Sns6/giphy-downsized-large.gif)

### Do-ocracy

Buster runs on do-ocracy. Do-ocracy is a simple concept. If you don't like something you don't need permission to fix it, you can just go ahead and fix it! If you actually want to merge your fix, or contribute in someway that benefits everybody, it'd really, really, really help if you got some light consensus from the rest of the Buster development community but hey, if you really need to do something then you just gotta do it! Just don't expect me to merge it if it doesn't meet our technical criteria or isn't quite right for Buster.

### Technical requirements

Part of what makes Buster so special is that we have standards. YAML based CICD hurts enough already and we just don't need to add to that.

All successfully merged pull requests must meet the following criteria: 

* All current tests must pass.
 
* At least one new test must be written to prove that the merged feature works correctly.

* At least one new [example test](https://blog.golang.org/examples) must be written to demonstrate the merged feature in our docs.
  
* Build tests must pass for all currently supported systems and package managers. Linux, Mac OSX, Windows, etc.
  
* Code must be clean, readable, and commented. How you do that is up to you!

Don't worry if you submit a pull request and all the tests break and the code is not readable. We won't merge it just yet and then you can get some feedback about what needs to be changed before we do!

### Be welcoming

As one final guideline please be welcoming to newcomers and encourage new contributors from all walks of life. I want Buster to be for everyone and that includes you and people who don't look, sound, or act like you!

# Your first contribution

Unsure where to begin contributing to Buster? You can start by looking through these beginner and help-wanted issues:

[Beginner issues](https://github.com/TimothyStiles/Buster/issues?q=is%3Aissue+is%3Aopen+label%3A%22beginner%22+) - issues which should only require a few lines of code, and a test or two.

[Good first issues](https://github.com/TimothyStiles/Buster/contribute) - issues which are good for first time contributors.

[Help wanted issues](https://github.com/TimothyStiles/Buster/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22+) - issues which should be a bit more involved than beginner issues.

[Feature requests](https://github.com/TimothyStiles/Buster/labels/enhancement) - before requesting a new feature search through previous feature requests to see if it's already been requested. If not then feel free to submit a request and tag it with the enhancement tag!

### Working on your first Pull Request? 

You can learn how from this *free* series, [How to Contribute to an Open Source Project on GitHub](https://egghead.io/series/how-to-contribute-to-an-open-source-project-on-github).

You can also check out [these](http://makeapullrequest.com/) [tutorials](http://www.firsttimersonly.com/).

At this point, you're ready to make your changes! Feel free to ask for help; everyone is a beginner at first :smile_cat:

# Getting started

For something that is bigger than a one or two line fix:

1. Create your own fork of the code.
2. Make a branch in your fork
3. Do the changes in your fork's branch.
4. Send a pull request.

# How to report a bug

### Security disclosures

If you find a security vulnerability, do NOT open an issue. I've yet to set up a security email for this so please in the interim DM me on twitter for my email [@timothystiles](https://twitter.com/TimothyStiles).

In order to determine whether you are dealing with a security issue, ask yourself these two questions:

* Can I access something that's not mine, or something I shouldn't have access to?
* Can I disable something for other people?
  
If the answer to either of those two questions are "yes", then you're probably dealing with a security issue. Note that even if you answer "no" to both questions, you may still be dealing with a security issue, so if you're unsure, just DM me [@timothystiles](https://twitter.com/TimothyStiles) for my personal email until I can set up a security related email.

### Non-security related bugs

For non-security bug reports please [submit it using this template!](https://github.com/TimothyStiles/Buster/issues/new?assignees=&labels=&template=bug_report.md&title=)

# How to suggest a feature or enhancement

If you want to suggest a feature it's as easy as filling out this [issue template](https://github.com/TimothyStiles/Buster/issues/new?assignees=&labels=&template=feature_request.md&title=), but before you do please [check to see if it's already been suggested!](https://github.com/TimothyStiles/Buster/labels/enhancement)

# In closing

Thanks, for reading and I'm super psyched to see what you'll do with Buster!
