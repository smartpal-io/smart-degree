import expectThrow from "../node_modules/openzeppelin-solidity/test/helpers/expectThrow";
import expectEvent from "../node_modules/openzeppelin-solidity/test/helpers/expectEvent";
import assertRevert from '../node_modules/openzeppelin-solidity/test/helpers/assertRevert';

const SmartDegreeMock = artifacts.require('SmartDegree');

require('chai')
  .use(require('chai-as-promised'))
.should();

contract('SmartDegree', function (accounts) {

  let mock;

  const gokuDegreeId = "0x000000000000000000000000000000000000000000000000000000000000001";
  const gokuDegreeHash = "0xe7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c";
  const vegetaDegreeId = "0x000000000000000000000000000000000000000000000000000000000000002";
  const vegetaDegreeHash = "0x46d0df9c7e28e2695adf3712936dcca8227b808060d28d81bd112ffb348fc3e4";
  const gohanDegreeId = "0x000000000000000000000000000000000000000000000000000000000000003";
  const gohanDegreeHash = "0x5cd27ecbeb083f226affb83b68806c385003e20be3fa663800ba81952ea31671";


  const [
    owner,
    anyone,
  ] = accounts;

  before(async function () {
    mock = await SmartDegreeMock.new();
  });

  context('in normal conditions', () => {
    it('should deliver the degree', async function () {
      await expectEvent.inTransaction(
        mock.deliverDegree(gokuDegreeId, gokuDegreeHash, { from: owner }),
        'LogDegreeDelivered'
      );
      const retrievedDegreeHash = await mock.getDegreeHash(gokuDegreeId);
      retrievedDegreeHash.should.be.equal(gokuDegreeHash);
    });

    it('should validate the degree', async function () {
          await expectEvent.inTransaction(
            mock.deliverDegree(vegetaDegreeId, vegetaDegreeHash, { from: owner }),
            'LogDegreeDelivered'
          );
          const isValid = await mock.isValid(vegetaDegreeId, vegetaDegreeHash);
          isValid.should.be.equal(true);
    });

    it('should not validate the degree because it was never delivered', async function () {
          const isValid = await mock.isValid(gohanDegreeId, gohanDegreeHash);
          isValid.should.be.equal(false);
    });

    it('should not delever the degree if the sender is not whitelisted', async function () {
      await assertRevert(
        mock.deliverDegree(gokuDegreeId, gokuDegreeHash, { from: anyone })
      );
    });

  });
});