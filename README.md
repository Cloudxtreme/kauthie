# Kauthie

Kauthie is a account/user managment implementation so I don't ever have to write that again.

I implemented this all login/logout/reset/invite/edit thing a few times and I am getting tiered of it,
I looked at *UserApp* and *StormPath* but I never was stoked @ outsourcing or depending on someone else
for a core part of an app.

This implementation is written on Go with the Gin framework and exposes a simple JSON API to allow other
parts of the app to interact with account data. OAuth2 is also implemented.

![Katie](https://raw.githubusercontent.com/kiasaki/kauthie/master/images/katie.png)
[Drawing](http://yassui.deviantart.com/art/Katie-301245033) by **yassui** @ DeviantArt

## Deploying / Demo

Setting up **Kauthie** on [Heroku](https://heroku.com/) is pretty straight foward:

```
git clone git@github.com:kiasaki/kauthie.git
hk create
hk addon-add mongohq
hk addon-add mandrill
hk set BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git
git push heroku master
```

## Technology

**Kauthie** is written in _Go_ using _MongoDB_ as datastore and _Stripe_ for payment processing and 
recurring billing.

This technology stack has a nice balance of _speed_, _maintainability_, _ease of deployment_ and 
_flexibility to adapt_ to different requirements.

The important libraries **Kauthie** relies upon are the following:

- [github.com/GeertJohan/go.rice](http://godoc.org/github.com/GeertJohan/go.rice)
- [github.com/gorilla/mux](http://godoc.org/github.com/gorilla/mux)
- [github.com/gorilla/sessions](http://godoc.org/github.com/gorilla/sessions)
- [gopkg.in/mgo.v2](http://godoc.org/gopkg.in/mgo.v2)
- [github.com/stripe/stripe](http://godoc.org/github.com/stripe/stripe)
- [github.com/bluele/gforms](http://godoc.org/github.com/bluele/gforms)
- [code.google.com/p/go.crypto/bcrypt](http://godoc.org/code.google.com/p/go.crypto/bcrypt)

## What's implemented?

- Signed out
  - [x] GET    _/signup_          => Signup form with CC info & account name & user info
  - [x] POST   _/signup_          => Creates account & Creates user & Links user to account & Creates Stripe custumer with selected plan
  - [x] GET    _/login_           => Login page, support "?next=http://app.exemple.com"
  - [x] POST   _/login_           => Redirects to next url or to "/"
  - [ ] GET    _/forgot_          => Asks for email to whom send recovery instructions
  - [ ] POST   _/forgot_          => Sends an email containg the nessesary information to change password
  - [ ] GET    _/change-password_ => Allows for changing password provided you have a key (email links send here)
  - [ ] GET    _/logout_          => Well ... logs you out!
- Signed in
  - [ ] GET    _/_                    => Redirects you to your first account or make your select one if you have multiple
  - [ ] GET    _/a/{id}/profile_      => Editing your personnal info (email, name, password)
  - [ ] POST   _/a/{id}/profile_      => Saves profile
  - [ ] GET    _/a/{id}/settings_     => Editing the accounts settings (account name, owner, delete account)
  - [ ] POST   _/a/{id}/settings_     => Saves account / Changes owner
  - [ ] DELETE _/a/{id}_              => Queues account for deletion
  - [ ] GET    _/a/{id}/billing_      => Plan seelction and billing details
  - [ ] POST   _/a/{id}/billing_      => Saves billing info & handles changing plan
  - [ ] GET    _/a/{id}/history_      => Show billing history (past bills)
  - [ ] GET    _/a/{id}/bill/{id}_    => Show bill details and offer print option
  - [ ] GET    _/a/{id}/users_        => Show account users list
  - [ ] GET    _/a/{id}/users/invite_ => Show account user invitation form
  - [ ] GET    _/authorize_           => OAuth 2.0 Authrization page (send to login before if not connected)
  - [ ] GET    _/token_               => OAuth 2.0 Token page, used for access tokens/refresh token validation during OAuth auth

## Contributing