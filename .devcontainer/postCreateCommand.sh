#!/bin/bash

lefthook install

# mcp install
# TODO: remove if not using Claude Code

# add only if context7 does not already exist
if ! claude mcp list 2>/dev/null | grep -q "context7"; then
    echo "Adding MCP server context7..."
    claude mcp add --transport http context7 https://mcp.context7.com/mcp
else
    echo "MCP server context7 already exists, skipping..."
fi

# add only if cipher does not already exist
if ! claude mcp list 2>/dev/null | grep -q "cipher"; then
    # check for API key
    if [ -z "$ANTHROPIC_API_KEY" ]; then
        echo "⚠️  WARNING: No API keys found for cipher MCP server."
        echo "   cipher requires: ANTHROPIC_API_KEY"
        echo "   Set the API key in your host environment (~/.zshrc or ~/.bashrc) and rebuild the DevContainer."
        echo "   Skipping cipher MCP server installation..."
    else
        echo "Adding MCP server cipher..."
        claude mcp add --transport stdio cipher -- cipher --mode mcp
    fi
else
    echo "MCP server cipher already exists, skipping..."
fi

# add only if serena does not already exist
if ! claude mcp list 2>/dev/null | grep -q "serena"; then
    echo "Adding MCP server serena..."
    claude mcp add serena -- uvx --from git+https://github.com/oraios/serena serena start-mcp-server \
        --context ide-assistant \
        --project "$(pwd)" \
        --enable-web-dashboard false \
        --enable-gui-log-window false
else
    echo "MCP server serena already exists, skipping..."
fi
