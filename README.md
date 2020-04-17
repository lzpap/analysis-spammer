# analysis-spammer
Utility to spam GoShimmer Analysis Server with autopeering heartbeat messages

# Prerequisites

```
pip install networkx
```

# Usage

1. Run the graph generator python script. This will generate graph data points:
   ```
   python graph_generator.py <number-of-nodes>
   ```
   The generated graph is an approximation of how nodes do autopeering in GoShimmer.

2. Run `main.go`. This will consume the generated data, and start sending heartbeat
   messages to `127.0.0.1:188`, the assumed address of the Analysis Server.
   ```
   go run main.go --patern=<flood, floodReverse, distribute> --nps=<number>
   ```

 For each generated node, the spammer sends one heartbeat message with all neighbors per 5 seconds to the server.
