using System;
using System.Collections.Generic;
using ApiFarm.Models.Impl;
using ApiFarm.Utils.Impl;
using NUnit.Framework;
using Shouldly;

namespace ApiFarm.Test.Utils
{
    public class VideoGameClonerTests
    {
        private VideoGameCloner subject;

        [SetUp]
        public void SetUp()
        {
            subject = new VideoGameCloner();
        }

        private class CloneShould : VideoGameClonerTests
        {
            [Test]
            public void ReturnDeepCloneWithDifferentReferences()
            {
                // Arrange
                var original = new VideoGame
                {
                    Id = 5,
                    Name = "Vikings at Sea II",
                    Developers = new List<string> { "A" },
                    Publishers = new List<string> { "A", "B" },
                    Directors = new List<string> { "A" },
                    Producers = new List<string> { "A" },
                    Designers = new List<string> { "A", "B", "C" },
                    Programmers = new List<string> { "A", "B", "C", "D" },
                    Artists = null,
                    Composers = new List<string> { "A", "B" },
                    Platforms = new List<string> { "A", "B", "C" },
                    DateReleased = DateTime.Now,
                };

                // Act
                var clone = subject.Clone(original);

                // Assert
                clone.ShouldBeEquivalentTo(original);

                clone.ShouldNotBeSameAs(original);
                clone.Developers.ShouldNotBeSameAs(original.Developers);
                clone.Publishers.ShouldNotBeSameAs(original.Publishers);
                clone.Directors.ShouldNotBeSameAs(original.Directors);
                clone.Producers.ShouldNotBeSameAs(original.Producers);
                clone.Designers.ShouldNotBeSameAs(original.Designers);
                clone.Programmers.ShouldNotBeSameAs(original.Programmers);
                clone.Composers.ShouldNotBeSameAs(original.Composers);
                clone.Platforms.ShouldNotBeSameAs(original.Platforms);
                clone.DateReleased.ShouldNotBeSameAs(original.DateReleased);
            }
        }
    }
}
