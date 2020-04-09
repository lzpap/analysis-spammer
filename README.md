# analysis-spammer
Utility to spam GoShimmer Analysis Server with autopeering heratbeat messages

# Prerequisites

```
pip install networkx
```

# Usage

1. Run the graph generator python script. This will generate graph data points:
   ```
   python graph_generator <number-of-nodes>
   ```
   The generated graph is an Erdős–Rényi random graph. On average, nodes should have 8 neighbors.
   The script will create two text files with data: `generated_links.txt` and `generated_nodes.txt`.

2. Run `main.go`. This will consume the generated data, and start sending heartbeat
   messages to `127.0.0.1:188`, the assumed address of the Analysis Server.

 For each generated node, the spammer sends one heartbeat message with all neighbors per 5 seconds to the server.