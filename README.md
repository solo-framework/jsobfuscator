
## Introduction

This is Go app that helps obfuscate JavaScript code by using JavaScript Obfuscator and V8 Engine.

## Requirements

- Go 1.24.0 or later
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

There are 3 levels of obfuscation: **low**, **medium**, **high**

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

## Contributing

Contributions are welcome. Please feel free to submit a Pull Request or create an Issue on GitHub.

## License

This project is distributed under the MIT license.
