using System.Collections.Generic;
using ApiFarm.Models;
using ApiFarm.Repositories;
using ApiFarm.Utils;
using Moq;
using NUnit.Framework;
using Shouldly;

namespace ApiFarm.Test.Repositories
{
    public class InMemoryTests
    {
        private InMemory<Model> subject;

        private Mock<ICloner<Model>> mockCloner;
        private List<Model> models;

        [SetUp]
        public void Setup()
        {
            mockCloner = new Mock<ICloner<Model>>();
            models = new List<Model>();
            subject = new InMemory<Model>(mockCloner.Object, models);
        }

        public class Model : IModel
        {
            public uint Id { get; set; }
        }

        private class ResetShould : InMemoryTests
        {
            [Test]
            public void ClearModelsList()
            {
                // Arrange
                models.Add(new Model());
                models.Add(new Model());
                models.Add(new Model());

                // Act
                subject.Reset();

                // Assert
                models.ShouldBeEmpty();
            }
        }

        private class GetAllShould : InMemoryTests
        {
            [Test]
            public void ReturnEmptyModelsList()
            {
                // Arrange
                subject = new InMemory<Model>(mockCloner.Object);

                // Act
                var retrievedModels = subject.GetAll();

                // Assert
                retrievedModels.ShouldBeEmpty();
            }

            [Test]
            public void ReturnListOfClonedModels()
            {
                // Arrange
                var originalModel = new Model();
                models.Add(originalModel);
                var cloneModel = new Model();

                mockCloner.Setup(m => m.Clone(originalModel)).Returns(cloneModel);

                // Act
                var retrievedModels = subject.GetAll();

                // Assert
                retrievedModels.ShouldNotBeSameAs(models);
                retrievedModels.ShouldContain(cloneModel);
            }
        }

        private class AddShould : InMemoryTests
        {
            [Test]
            public void SetIncrementedIdOfModelAndAddToList()
            {
                // Arrange
                var originalModel1 = new Model();
                var cloneModel1 = new Model();
                var originalModel2 = new Model();
                var cloneModel2 = new Model();

                mockCloner.Setup(m => m.Clone(originalModel1)).Returns(cloneModel1);
                mockCloner.Setup(m => m.Clone(originalModel2)).Returns(cloneModel2);

                // Act
                var addedModel1 = subject.Add(originalModel1);
                var addedModel2 = subject.Add(originalModel2);

                // Assert
                models.ShouldContain(cloneModel1);
                models.ShouldContain(cloneModel2);

                addedModel1.ShouldBe(originalModel1);
                addedModel2.ShouldBe(originalModel2);

                addedModel1.Id.ShouldBe(1u);
                addedModel2.Id.ShouldBe(2u);
            }
        }
    }
}
