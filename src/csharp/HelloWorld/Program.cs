using System;
using System.Threading;

namespace HelloWorld
{
    class Program
    {
        static void Main(string[] args)
        {
            while (true)
            {
                // Reset cursor to beginning of line
                Console.SetCursorPosition(0, Console.CursorTop);

                var current = DateTime.Now;
                // Write used to stay on same line
                Console.Write(current.ToLongTimeString());
                var badmath = new BadMath();
                var x = 5;
                var xfact = badmath.factorial(5);

                Thread.Sleep(1000);
            }
        }
    }
}
