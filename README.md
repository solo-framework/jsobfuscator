
## Introduction

This is a Go app that helps obfuscate JavaScript code by using JavaScript Obfuscator and V8 Engine.

## Requirements

- Go 1.26.1 or later
- v8go library

## Usage

### Build for linux
```bash
make build
```

### Obfuscation

```bash
$jsobf -l medium < original.js > obfuscated.js
```

There are 3 levels of obfuscation: low, medium, high

### Important Notes

- JavaScript code cannot contain backtick characters (`) as they are used to wrap the code during processing.
- Javascript comments only as /** comment here  */
- The obfuscator uses the following default options:
  ```javascript
  const options = {
      compact: (Math.random() < 0.5),
      controlFlowFlattening: true,
      controlFlowFlatteningThreshold: 1,
      numbersToExpressions: true,
      simplify: true,
      stringArrayShuffle: true,
      splitStrings: true,
      stringArrayThreshold: 1
  }
  ```

## Performance Optimization

The library uses a caching mechanism to optimize performance when performing multiple obfuscations. You should reuse the same Obfuscator instance when you need to obfuscate multiple JavaScript code snippets.

```go
obf, _ := obfuscator.NewObfuscator()

// Use the same instance for multiple obfuscations
result1, _ := obf.Obfuscate(jsCode1)
result2, _ := obf.Obfuscate(jsCode2)
result3, _ := obf.Obfuscate(jsCode3)
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request or create an Issue on GitHub.

## License

This project is distributed under the MIT license.
