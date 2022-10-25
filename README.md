# graph_backend

Project: Software Engineering

Basic requirements for UI:

- setup for #nodes
- setup for edges: direction, weight
- separated checkboxes (or similar): analysis tool #1, analysis tool #2, etc...
- graph download option as PNG, PDF
- dynamic graph visualisation: each new edge should appear real time (so NOT all edges at once)

DEVELOPMENT DEADLINE: 12th Oct.

Analysis tools between node1 and node2: sum 7

- Generate all paths
- Generate paths with up to N edges
- Generate paths with up to W summed weight
- Generate path with lowest summed weight
- Generate path with highest summed weight
- Generate path with the fewest edges
- Generate path with the most edges

About path generation:

- no edge repetition is allowed in a path
- it returns path(s) in a vector-type data structure
- it returns it empty if no path found

Additional requirements:

- #nodes between 2 and 15 - no other number should be allowed in field
- minimum #edges = 1
- edge weight is always a positive floating number - no other value should be allowed in field
- Direction of an edge: if directed: A --> B direction, (not A <--- B), if not, this information has to be passed to backend and the visualised graph should use a line instead of an arrow.
- Default edge weight: 1.0. If weight is not typed in by the user, the program sets it to this default value.
- WE name the nodes, not the user. It has to follow simple alphabetical order (since max #nodes = 15, we'll have enough letters).
