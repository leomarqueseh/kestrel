# Contributing

This is currently a solo learning/portfolio project, but it follows the same conventions a real engineering team would use.

## Language

All code, comments, commit messages, branch names, API responses, and logs are written in **English**.

## Commit convention

[Conventional Commits](https://www.conventionalcommits.org/):

```
feat(api): add target management endpoints
fix(scanner): handle timeout on unreachable target
docs: update architecture diagram
refactor(auth): extract token validation into middleware
test(scanner): add reconnaissance module tests
security(api): enforce authorization middleware
build: bump go.mod to 1.23
ci: add container vulnerability scanning
chore: update dependencies
```

## Branch strategy

- `main` — always deployable
- `feature/<short-description>` — one phase or sub-feature per branch
- Merge via pull request, even solo — keeps a reviewable history

## Code standard

Code should be readable, idiomatic, testable, secure, modular, and documented. Comments explain **why**, not **what**:

```go
// Bad
// increment counter
counter++

// Good
// Increment the retry counter so transient failures can be handled
// without exceeding the configured limit.
counter++
```

## Definition of done (per phase)

A phase is complete only when:

- [ ] Code implemented
- [ ] Tests passing
- [ ] Security review done
- [ ] Documentation updated (including README when relevant)
- [ ] Git commit created
- [ ] Code reviewed
- [ ] Result validated against the phase's expected outcome
