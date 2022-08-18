# 1) Workspace nested modules
- go does not have real 'sub'-modules
- all modules are equal independent of the filesystem layout
- go commands take only effect in current module
  - ignores nested modules
- go workspace can be used to handle cross-module dependencies

# 2) Nested modules
- go does not have real 'sub'-modules
- all modules are equal independent of the filesystem layout
- go commands take only effect in current module
    - ignores nested modules

# 3) Single module
- no modules, no differentiation
- go commands take effect to everything in module