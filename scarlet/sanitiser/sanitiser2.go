// Sanitise package removes tokens redundant or inconvenient to parsing.
// Traditionally, this process is performed during the scanning process but by
// decoupling the sanitisation from scanning the scanner can be reused in
// source code formatting and analysis tools.
package sanitiser
