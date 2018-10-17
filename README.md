# Police UK API

A work in progress wrapper for the data.police.uk api

- [ ] Rate limit // _need tests to check rate limit does not exceed 15 requests per second_

## Forces

- [x] Forces
- [x] Specific force
- [x] Force senior officers 

## Crime

- [x] Street level crimes
- [x] Street level outcomes
- [x] Crimes at location
- [x] Crimes with no location
- [x] Crime categories
- [x] Last updated
- [x] Outcomes for a specific crime

## Neighbourhood related

- [x] Neighbourhoods
- [x] Specific neighbourhood
- [x] Neighbourhood boundary
- [x] Neighbourhood team
- [x] Neighbourhood events
- [x] Neighbourhood priorities

## Locate neighbourhood

- [x] Stop and searches by area
- [x] Stop and searches by location
- [x] Stop and searches with no location
- [x] Stop and searches by force

## Misc. ToDo

- [ ] @TODO: handle 503 error, which is returned when a request contains over 10000 results 
- [ ] Handle this case: The API will return a 400 status code in response to a GET request longer than 4094 characters. For submitting particularly complex poly parameters, consider using POST instead.
- [ ] Think on how to handle multiple type returns by the Stop and Search outcome variable.

