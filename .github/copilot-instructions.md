# Copilot Instructions

## Carbon View Components

When working on files in `pkg/carbon` that define or modify Carbon view components:

- Read `pkg/carbon/README.md` first and treat it as required reference material before proposing or making changes.
- Refer back to `pkg/carbon/README.md` whenever the work touches Carbon view construction, slots, attrs, component state, or events.
- Preserve the existing Carbon view structure: `View...` constants in `view.go`, concrete types embedding `base`, `init()` registration with `mvc.RegisterView(..., mvc.NewViewWithElement(..., setView))`, and exported constructors using `mvc.NewView(..., setView, args)`.
- Prefer the existing typed attribute helpers in `pkg/carbon/attr.go` instead of introducing ad hoc string attributes.
- Expose component state only through the relevant `pkg/mvc/state.go` interfaces and their standard getter/setter methods; do not introduce custom state-mutating APIs when an existing MVC contract applies.
- Use the event constants in `pkg/carbon/event.go` as the public event surface, and keep emitted event targets at the component boundary rather than exposing shadow-root implementation details.
- For simple components, use a single tag or Carbon custom element. For structured components, use template-backed views with `data-slot` markers.
- Update the matching component markdown doc in `pkg/carbon/*.md` when behavior or public constructor usage changes.
- Use `pkg/carbon/README.md` as the package-level source of truth for how Carbon views are created and extended.